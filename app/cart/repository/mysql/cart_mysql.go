package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
	"time"
)

type CartRepoMysql struct {
	DB *gorm.DB
}

func NewCartRepoMysql(db *gorm.DB) domain.CartRepository {
	return &CartRepoMysql{DB: db}
}

func (c *CartRepoMysql) AddProduct(ctx context.Context,
	email string,
	productID, quantity int) error {
	if err := c.DB.
		WithContext(ctx).
		Error; err != nil {
		return err
	}
	return nil
}

func (c *CartRepoMysql) Checkout(ctx context.Context,
	products []domain.Product,
	startAt, finishAt time.Time) error {
	if err := c.DB.
		WithContext(ctx).
		Error; err != nil {
		return err
	}
	return nil
}
