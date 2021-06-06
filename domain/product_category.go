package domain

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name     string    `gorm:"unique;not null" json:"name"`
	Products []Product `gorm:"foreignKey:ProductCategoryID" json:"products"`
}

type ProductCategoryRepository interface {
}

type ProductCategoryUseCase interface {
}
