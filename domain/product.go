package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name              string         `gorm:"not null" json:"name"`
	Price             uint           `gorm:"not null" json:"price"`
	Stock             uint           `gorm:"not null" json:"stock"`
	ProductCategoryID uint           `json:"product_category_id"`
	Images            []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
}
