package usecase

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"time"
)

type CartUseCase struct {
	cartRepoMysql domain.CartRepository
	userRepoMysql domain.UserRepository
	timeOut       time.Duration
}

func NewCartUseCase(c domain.CartRepository,
	u domain.UserRepository,
	timeOut time.Duration) domain.CartUseCase {
	return &CartUseCase{
		cartRepoMysql: c,
		userRepoMysql: u,
		timeOut:       timeOut,
	}
}

func (c *CartUseCase) AddProductToCart(ctx context.Context, productID int, email string, q *domain.CartAddProductRequest) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()

	user, err := c.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err := c.cartRepoMysql.AddProductToCart(ctx,
		q.Quantity,
		uint(productID),
		user.ID); err != nil {
		return err
	}

	return nil
}
