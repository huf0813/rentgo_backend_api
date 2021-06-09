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
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Price           uint    `json:"price"`
	Stock           uint    `json:"stock"`
	Star            float64 `json:"star"`
	Reviews         uint    `json:"reviews"`
	ProductCategory string  `json:"product_category"`
	Vendor          string  `json:"vendor"`
}

type ProductReviewResponse struct {
	UserName      string `json:"user_name"`
	ProductRating uint   `json:"product_rating"`
	ProductReview string `json:"product_review"`
}

type ProductImageResponse struct {
	Path string `json:"path"`
}

type ProductRepository interface {
	FetchByID(ctx context.Context, id int) (ProductResponse, error)
	FetchByCategory(ctx context.Context, category string) ([]ProductResponse, error)
	SearchProduct(ctx context.Context, name string) ([]ProductResponse, error)
	FetchReviewsByID(ctx context.Context, id int) ([]ProductReviewResponse, error)
	FetchImagesByID(ctx context.Context, id int) ([]ProductImageResponse, error)
}
type ProductUseCase interface {
	FetchByID(ctx context.Context, id int) (ProductResponse, error)
	FetchByCategory(ctx context.Context, category string) ([]ProductResponse, error)
	SearchProduct(ctx context.Context, name string) ([]ProductResponse, error)
	FetchReviewsByID(ctx context.Context, id int) ([]ProductReviewResponse, error)
	FetchImagesByID(ctx context.Context, id int) ([]ProductImageResponse, error)
}
