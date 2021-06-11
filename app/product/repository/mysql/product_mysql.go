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

func (p *ProductRepository) FetchImagesByID(ctx context.Context, id int) ([]domain.ProductImageResponse, error) {
	var result []domain.ProductImageResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select(
			"product_images.path as path, "+
				"product_images.id as product_image_id").
		Joins("JOIN product_images ON products.id = product_images.product_id").
		Where("products.id = ?", id).
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
			"users.name as vendor, " +
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, " +
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews," +
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Order("products.created_at").
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
			"users.name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, "+
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("products.id = ?", id).
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
			"users.name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, "+
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("product_categories.name LIKE ?", category+"%").
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
			"users.name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(ip.rating), 0), 1) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as star, "+
			"(SELECT COUNT(*) FROM invoice_products ip JOIN invoices i ON ip.invoice_id = i.id WHERE ip.product_id = products.id AND i.invoice_category_id = 2) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("products.name LIKE ?", name+"%").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
