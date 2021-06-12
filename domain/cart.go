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
	ID           int    `json:"cart_id"`
	Vendor       string `json:"product_vendor"`
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
	Quantity     uint   `json:"product_quantity"`
}

type CartAddProductRequest struct {
	Quantity int `json:"quantity"`
}

type CartRepository interface {
	AddProductToCart(ctx context.Context,
		quantity int,
		productID,
		userID uint) error
	FetchCart(ctx context.Context,
		userID uint) ([]CartResponse, error)
	FetchCartByID(ctx context.Context,
		userID, cartID uint) (Cart, error)
	DeleteCartByID(ctx context.Context, userID, cartID uint) error
	IsCartByIDsExist(ctx context.Context, userID uint, cartIDs []uint) (bool, error)
}

type CartUseCase interface {
	AddProductToCart(ctx context.Context,
		productID int,
		email string,
		q *CartAddProductRequest) error
	FetchCart(ctx context.Context, email string) ([]CartResponse, error)
	DeleteCartByID(ctx context.Context, email string, cartID uint) error
}
