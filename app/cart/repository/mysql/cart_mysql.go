package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type CartRepoMysql struct {
	DB *gorm.DB
}

func NewCartRepoMysql(db *gorm.DB) domain.CartRepository {
	return &CartRepoMysql{DB: db}
}

func (c *CartRepoMysql) AddProductToCart(ctx context.Context,
	quantity int,
	productID,
	userID uint) error {
	newCart := domain.Cart{
		ProductID: productID,
		UserID:    userID,
		Quantity:  uint(quantity),
	}

	if err := c.DB.
		WithContext(ctx).
		Create(&newCart).Error; err != nil {
		return err
	}

	return nil
}
