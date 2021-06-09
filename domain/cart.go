package domain

import (
	"context"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	UserID    uint `json:"user_id"`
	Quantity  uint `json:"quantity"`
}

type CartResponse struct {
	ProductName string `json:"product_name"`
	Quantity    uint   `json:"quantity"`
}

type CartAddProductRequest struct {
	Quantity int `json:"quantity"`
}

type CartRepository interface {
	AddProductToCart(ctx context.Context, quantity int, productID, userID uint) error
}

type CartUseCase interface {
	AddProductToCart(ctx context.Context, productID int, email string, q *CartAddProductRequest) error
}
