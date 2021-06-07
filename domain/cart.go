package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	ProductID  uint      `json:"product_id"`
	UserID     uint      `json:"user_id"`
	Quantity   uint      `json:"quantity"`
	StartDate  time.Time `json:"start_date"`
	FinishDate time.Time `json:"finish_date"`
}

type CartProductUpdateDateRequest struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

type CartUpdateDateRequest struct {
	StartAt      string                         `json:"start_at"`
	FinishAt     string                         `json:"finish_at"`
	CartProducts []CartProductUpdateDateRequest `json:"cart_products"`
}

type CartResponse struct {
	ProductName string `json:"product_name"`
	Quantity    uint   `json:"quantity"`
}

type CartUseCase interface {
	FetchCart(ctx context.Context, email string) error
	AddProduct(ctx context.Context, email string, productID, quantity int) error
	Checkout(ctx context.Context, products []Product, startAt, finishAt time.Time) error
}

type CartRepository interface {
	AddProduct(ctx context.Context, email string, productID, quantity int) error
	Checkout(ctx context.Context, products []Product, startAt, finishAt time.Time) error
}
