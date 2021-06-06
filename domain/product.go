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
	Star              uint             `gorm:"not null" json:"star"`
	ProductCategoryID uint             `json:"product_category_id"`
	ProductImages     []ProductImage   `gorm:"foreignKey:ProductID" json:"product_images"`
	InvoiceReviews    []InvoiceProduct `gorm:"foreignKey:ProductID" json:"invoice_reviews"`
	EventProducts     []EventProduct   `gorm:"foreignKey:ProductID" json:"event_products"`
}

type ProductResponse struct {
	Name            string        `json:"name"`
	Price           uint          `json:"price"`
	Stock           uint          `json:"stock"`
	Star            uint          `json:"star"`
	ProductCategory string        `json:"product_category"`
	Images          []interface{} `json:"images"`
	Reviews         []interface{} `json:"reviews"`
}

type ProductRepository interface {
	FetchByID(ctx context.Context, id int) (ProductResponse, error)
	FetchByCategory(ctx context.Context, category string) ([]ProductResponse, error)
	SearchProduct(ctx context.Context, name string) ([]ProductResponse, error)
}
type ProductUseCase interface {
	FetchByID(ctx context.Context, id int) (ProductResponse, error)
	FetchByCategory(ctx context.Context, category string) ([]ProductResponse, error)
	SearchProduct(ctx context.Context, name string) ([]ProductResponse, error)
}
