package domain

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name              string           `gorm:"not null" json:"name"`
	Price             uint             `gorm:"not null" json:"price"`
	Stock             uint             `gorm:"not null" json:"stock"`
	ProductCategoryID uint             `json:"product_category_id"`
	Images            []ProductImage   `gorm:"foreignKey:ProductID" json:"images"`
	Reviews           []InvoiceProduct `gorm:"foreignKey:ProductID" json:"reviews"`
}

type ProductRepository interface {
	Fetch(ctx context.Context)
	FetchByID(ctx context.Context, id int)
	FetchByCategory(ctx context.Context, category string)
	SearchProduct(ctx context.Context, name string)
}

type ProductUseCase interface {
	Fetch(ctx context.Context)
	FetchByID(ctx context.Context, id int)
	FetchByCategory(ctx context.Context, category string)
	SearchProduct(ctx context.Context, name string)
}
