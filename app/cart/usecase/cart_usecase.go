package usecase

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"time"
)

type CartUseCase struct {
	cartRepoMysql domain.CartRepository
	timeOut       time.Duration
}

func NewCartUseCase(c domain.CartRepository, timeOut time.Duration) domain.CartUseCase {
	return &CartUseCase{
		cartRepoMysql: c,
		timeOut:       timeOut,
	}
}

func (c *CartUseCase) FetchCart(ctx context.Context, email string) error {
	panic("implement me")
}

func (c *CartUseCase) AddProduct(ctx context.Context, email string, productID, quantity int) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()

	if err := c.cartRepoMysql.AddProduct(ctx,
		email,
		productID,
		quantity); err != nil {
		return err
	}

	return nil
}

func (c *CartUseCase) Checkout(ctx context.Context, products []domain.Product, startAt, finishAt time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()

	if err := c.cartRepoMysql.Checkout(ctx,
		products,
		startAt,
		finishAt); err != nil {
		return err
	}

	return nil
}
