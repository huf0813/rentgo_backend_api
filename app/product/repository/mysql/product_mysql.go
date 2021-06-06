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

func (p *ProductRepository) FetchByID(ctx context.Context, id int) (domain.ProductResponse, error) {
	var product domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("product.name, "+
			"product.price, "+
			"product.stock, "+
			"product.star, "+
			"product_category.name as product_category").
		First(&product, id).
		Error; err != nil {
		return domain.ProductResponse{}, err
	}
	return product, nil
}

func (p *ProductRepository) FetchByCategory(ctx context.Context, category string) ([]domain.ProductResponse, error) {
	var result []domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("product.name, "+
			"product.price, "+
			"product.stock, "+
			"product.star, "+
			"product_category.name as product_category").
		Where("product_category.name LIKE ?", category+"%").
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductRepository) SearchProduct(ctx context.Context, name string) ([]domain.ProductResponse, error) {
	var result []domain.ProductResponse
	if err := p.DB.
		WithContext(ctx).
		Model(&domain.Product{}).
		Select("product.name, "+
			"product.price, "+
			"product.stock, "+
			"product.star, "+
			"product_category.name as product_category").
		Where("product.name LIKE ?", name+"%").Error; err != nil {
		return nil, err
	}
	return result, nil
}
