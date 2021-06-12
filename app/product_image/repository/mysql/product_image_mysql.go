package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type ProductImageRepoMysql struct {
	DB *gorm.DB
}

func NewProductImageRepoMysql(db *gorm.DB) domain.ProductImageRepository {
	return &ProductImageRepoMysql{DB: db}
}

func (p *ProductImageRepoMysql) CreateProductImage(ctx context.Context,
	productID uint,
	filename string) error {
	create := domain.ProductImage{
		Path:      filename,
		ProductID: productID,
	}

	if err := p.DB.WithContext(ctx).Create(&create).Error; err != nil {
		return err
	}

	return nil
}
