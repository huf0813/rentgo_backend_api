package domain

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name     string    `gorm:"unique;not null" json:"name"`
	Products []Product `gorm:"foreignKey:ProductCategoryID" json:"products"`
}

type ProductCategoryResponse struct {
	ProductCategoryID   uint   `json:"product_category_id"`
	ProductCategoryName string `json:"product_category_name"`
}
