package domain

import (
	"context"
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	Path      string `gorm:"unique;not null" json:"path"`
	ProductID uint   `json:"product_id"`
}

type ProductImageRepository interface {
	CreateProductImage(ctx context.Context,
		productID uint,
		filename string) error
}

type ProductImageUseCase interface {
}
