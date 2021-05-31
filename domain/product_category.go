package domain

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:ProductCategoryID" json:"products"`
}
