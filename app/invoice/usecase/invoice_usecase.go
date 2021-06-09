package usecase

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"time"
)

type InvoiceUseCase struct {
	invoiceRepoMysql domain.InvoiceRepository
	cartRepoMysql    domain.CartRepository
	userRepoMysql    domain.UserRepository
	timeOut          time.Duration
}

func NewInvoiceUseCase(i domain.InvoiceRepository,
	c domain.CartRepository,
	u domain.UserRepository,
	timeOut time.Duration) domain.InvoiceUseCase {
	return &InvoiceUseCase{
		invoiceRepoMysql: i,
		cartRepoMysql:    c,
		userRepoMysql:    u,
		timeOut:          timeOut,
	}
}

func (i *InvoiceUseCase) CreateCheckOut(ctx context.Context,
	startDate, finishDate time.Time,
	email string,
	cartIDS []int) error {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(
		ctx, email)
	if err != nil {
		return err
	}

	var invoiceProducts []domain.Cart
	for _, v := range cartIDS {
		res, err := i.cartRepoMysql.FetchCartByID(ctx,
			uint(v),
			user.ID)
		if err != nil {
			return err
		}
		invoiceProducts = append(invoiceProducts, res)
	}

	if err := i.invoiceRepoMysql.CreateCheckOut(ctx,
		startDate,
		finishDate,
		user.ID,
		invoiceProducts); err != nil {
		return err
	}

	return nil
}
