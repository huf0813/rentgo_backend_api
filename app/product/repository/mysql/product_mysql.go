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

func (p *ProductRepository) FetchReviewsByID(ctx context.Context, id int) ([]domain.ProductReviewResponse, error) {
	var result []domain.ProductReviewResponse
	if err := p.DB.
		WithContext(ctx).
		Table("product").
		Select(
			"users.name as user_name, "+
				"invoices_product.review as product_review, "+
				"invoices_product.rating as product_rating").
		Joins("JOIN invoices_product ON products.product_id = products.id").
		Joins("JOIN invoices ON invoices_product.invoice_id = invoices.id").
		Joins("JOIN users ON invoices.user_id = users.id").
		Where("products.id = ?", id).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductRepository) FetchByID(ctx context.Context, id int) (domain.ProductResponse, error) {
	var product domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Table("products").
		Preload("ProductImages").
		Select("products.name, "+
			"products.price, "+
			"products.stock, "+
			"products.id, "+
			"users.name as vendor, "+
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
		Table("products").
		Preload("ProductImages").
		Select("products.name, "+
			"products.price, "+
			"products.stock, "+
			"products.id, "+
			"users.name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(invoice_products.rating), 0), 1) FROM invoice_products WHERE invoice_products.product_id = products.id) as star, "+
			"(SELECT COUNT(*) FROM invoice_products WHERE invoice_products.product_id = products.id) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("product_categories.name LIKE ?", category+"%").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductRepository) SearchProduct(ctx context.Context, name string) ([]domain.ProductResponse, error) {
	var result []domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Table("products").
		Preload("ProductImages").
		Select("products.name, "+
			"products.price, "+
			"products.stock, "+
			"products.id, "+
			"users.name as vendor, "+
			"(SELECT ROUND(IFNULL(AVG(invoice_products.rating), 0), 1) FROM invoice_products WHERE invoice_products.product_id = products.id) as star, "+
			"(SELECT COUNT(*) FROM invoice_products WHERE invoice_products.product_id = products.id) as reviews,"+
			"product_categories.name as product_category").
		Joins("JOIN product_categories ON products.product_category_id = product_categories.id").
		Joins("JOIN users ON products.user_id = users.id").
		Where("products.name LIKE ?", name+"%").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
