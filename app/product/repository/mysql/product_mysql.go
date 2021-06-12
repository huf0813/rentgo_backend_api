package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &ProductRepository{DB: db}
}

func (p *ProductRepository) VendorFetchProductCategory(ctx context.Context) ([]domain.ProductCategoryResponse, error) {
	var res []domain.ProductCategoryResponse

	if err := p.DB.
		WithContext(ctx).
		Model(&domain.ProductCategory{}).
		Select("product_categories.name as product_category_name, " +
			"product_categories.id as product_category_id").
		Find(&res).
		Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (p *ProductRepository) VendorCreateProduct(ctx context.Context, pc *domain.Product) (uint, error) {
	if err := p.DB.
		WithContext(ctx).
		Create(&pc).
		Error; err != nil {
		return 0, err
	}
	return pc.ID, nil
}

func (p *ProductRepository) FetchImagesByID(ctx context.Context, id int) ([]domain.ProductImageResponse, error) {
	var result []domain.ProductImageResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select(
			"product_images.id, "+
				"product_images.path").
		Joins("JOIN users ON users.id = products.user_id").
		Joins("RIGHT JOIN product_images ON products.id = product_images.product_id").
		Where("product_images.product_id = ?", id).
		Where("users.identity_image IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_number IS NOT NULL").
		Where("users.store_name IS NOT NULL").
		Where("users.phone IS NOT NULL").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductRepository) FetchReviewsByID(ctx context.Context, id int) ([]domain.ProductReviewResponse, error) {
	var result []domain.ProductReviewResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select(
			"invoice_products.id as product_review_id, "+
				"users.name as user_name, "+
				"invoice_products.review as product_review, "+
				"invoice_products.rating as product_rating").
		Joins("JOIN invoice_products ON invoice_products.product_id = products.id").
		Joins("JOIN invoices ON invoice_products.invoice_id = invoices.id").
		Joins("JOIN users ON invoices.user_id = users.id").
		Where("products.id = ?", id).
		Where("invoices.invoice_category_id = ?", 2).
		Where("users.identity_image IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_number IS NOT NULL").
		Where("users.store_name IS NOT NULL").
		Where("users.phone IS NOT NULL").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductRepository) FetchLatestProduct(ctx context.Context) ([]domain.ProductResponse, error) {
	var result []domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("products.name, " +
			"products.price, " +
			"products.overview, " +
			"products.stock, " +
			"products.id, " +
			"users.store_name as vendor, " +
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, " +
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews," +
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("users.identity_image IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_number IS NOT NULL").
		Where("users.store_name IS NOT NULL").
		Where("users.phone IS NOT NULL").
		Order("products.created_at").
		Limit(5).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductRepository) FetchTrendingProduct(ctx context.Context) ([]domain.ProductResponse, error) {
	var result []domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("products.name, " +
			"products.price, " +
			"products.overview, " +
			"products.stock, " +
			"products.id, " +
			"users.store_name as vendor, " +
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, " +
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews," +
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("users.identity_image IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_number IS NOT NULL").
		Where("users.store_name IS NOT NULL").
		Where("users.phone IS NOT NULL").
		Order("star desc").
		Order("reviews desc").
		Limit(5).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductRepository) FetchByID(ctx context.Context, id int) (domain.ProductResponse, error) {
	var product domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("products.name, "+
			"products.price, "+
			"products.overview, "+
			"products.stock, "+
			"products.id, "+
			"users.store_name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, "+
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("products.id = ?", id).
		Where("users.identity_image IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_number IS NOT NULL").
		Where("users.store_name IS NOT NULL").
		Where("users.phone IS NOT NULL").
		First(&product).Error; err != nil {
		return domain.ProductResponse{}, err
	}
	return product, nil
}

func (p *ProductRepository) FetchByCategory(ctx context.Context, category string) ([]domain.ProductResponse, error) {
	var result []domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("products.name, "+
			"products.price, "+
			"products.overview, "+
			"products.stock, "+
			"products.id, "+
			"users.store_name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, "+
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("product_categories.name LIKE ?", category+"%").
		Where("users.identity_image IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_number IS NOT NULL").
		Where("users.store_name IS NOT NULL").
		Where("users.phone IS NOT NULL").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductRepository) SearchProduct(ctx context.Context, name string) ([]domain.ProductResponse, error) {
	var result []domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("products.name, "+
			"products.price, "+
			"products.overview, "+
			"products.stock, "+
			"products.id, "+
			"users.store_name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, "+
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("products.name LIKE ?", name+"%").
		Where("users.identity_image IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_number IS NOT NULL").
		Where("users.store_name IS NOT NULL").
		Where("users.phone IS NOT NULL").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
