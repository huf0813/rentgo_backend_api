package domain

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	Path      string `gorm:"not null" json:"path"`
	ProductID uint   `json:"product_id"`
}

type ProductImageRepository interface {
}

type ProductImageUseCase interface {
}
