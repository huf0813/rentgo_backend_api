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
	UserID            uint             `json:"user_id"`
	ProductImages     []ProductImage   `gorm:"foreignKey:ProductID" json:"product_images"`
	InvoiceProducts   []InvoiceProduct `gorm:"foreignKey:ProductID" json:"invoice_reviews"`
	EventProducts     []EventProduct   `gorm:"foreignKey:ProductID" json:"event_products"`
	Carts             []Cart           `gorm:"foreignKey:ProductID" json:"carts"`
}

type ProductResponse struct {
	ID              uint           `json:"id"`
	Name            string         `json:"name"`
	Price           uint           `json:"price"`
	Stock           uint           `json:"stock"`
	Star            float64        `json:"star"`
	Reviews         uint           `json:"reviews"`
	ProductCategory string         `json:"product_category"`
	Vendor          string         `json:"vendor"`
	ProductImages   []ProductImage `gorm:"foreignKey:ProductID" json:"product_images"`
}

type ProductDetailResponse struct {
	ID              uint             `json:"id"`
	Name            string           `json:"name"`
	Price           uint             `json:"price"`
	Stock           uint             `json:"stock"`
	Star            float64          `json:"star"`
	Reviews         uint             `json:"reviews"`
	ProductCategory string           `json:"product_category"`
	Vendor          string           `json:"vendor"`
	ProductImages   []ProductImage   `gorm:"foreignKey:ProductID" json:"product_images"`
	ProductReviews  []InvoiceProduct `gorm:"foreignKey:ProductID" json:"product_reviews"`
}

type ProductRepository interface {
	FetchByID(ctx context.Context, id int) (ProductDetailResponse, error)
	FetchByCategory(ctx context.Context, category string) ([]ProductResponse, error)
	SearchProduct(ctx context.Context, name string) ([]ProductResponse, error)
}
type ProductUseCase interface {
	FetchByID(ctx context.Context, id int) (ProductDetailResponse, error)
	FetchByCategory(ctx context.Context, category string) ([]ProductResponse, error)
	SearchProduct(ctx context.Context, name string) ([]ProductResponse, error)
}
