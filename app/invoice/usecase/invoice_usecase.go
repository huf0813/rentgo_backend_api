package usecase

import (
	"context"
	"errors"
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

func (i *InvoiceUseCase) CreateReviewByInvoiceProductID(ctx context.Context,
	invoiceProductID, star uint,
	review string) error {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	if err := i.invoiceRepoMysql.CreateReviewByInvoiceProductID(ctx,
		invoiceProductID, star, review); err != nil {
		return err
	}

	return nil
}

func (i *InvoiceUseCase) GetInvoicesAccepted(ctx context.Context,
	email string) ([]domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	res, err := i.invoiceRepoMysql.GetInvoiceByCategory(ctx,
		user.ID, 3)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *InvoiceUseCase) GetInvoicesOnGoing(ctx context.Context,
	email string) ([]domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	res, err := i.invoiceRepoMysql.GetInvoiceByCategory(ctx,
		user.ID, 1)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *InvoiceUseCase) GetInvoicesCompleted(ctx context.Context,
	email string) ([]domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	res, err := i.invoiceRepoMysql.GetInvoiceByCategory(ctx,
		user.ID, 2)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *InvoiceUseCase) GetInvoiceProductsByReceiptCode(ctx context.Context,
	email, receiptCode string) ([]domain.InvoiceProductResponse,
	error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	res, err := i.invoiceRepoMysql.GetInvoiceProductByReceiptNumber(ctx,
		user.ID,
		receiptCode)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *InvoiceUseCase) CreateCheckOut(ctx context.Context,
	startDate, finishDate time.Time,
	email string,
	cartIDS []uint) error {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(
		ctx, email)
	if err != nil {
		return err
	}

	isCartIDsValid, err := i.cartRepoMysql.IsCartByIDsExist(ctx,
		user.ID, cartIDS)
	if err != nil {
		return err
	}
	if !isCartIDsValid {
		return errors.New("one of cart_ids is invalid")
	}

	var invoiceProducts []domain.Cart
	for k := 0; k < len(cartIDS); k++ {
		res, err := i.cartRepoMysql.FetchCartByID(ctx,
			user.ID,
			cartIDS[k])
		if err != nil {
			return err
		}
		invoiceProducts = append(invoiceProducts, res)
		if err := i.cartRepoMysql.DeleteCartByID(ctx,
			user.ID,
			cartIDS[k]); err != nil {
			return err
		}
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

func (i *InvoiceUseCase) UpdateInvoiceOnGoing(ctx context.Context,
	email,
	receiptCode string) error {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err := i.
		invoiceRepoMysql.
		UpdateInvoiceCategory(
			ctx,
			user.ID,
			receiptCode,
			1); err != nil {
		return err
	}

	return nil
}

func (i *InvoiceUseCase) UpdateInvoiceCompleted(ctx context.Context,
	email string,
	receiptCode string) error {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err := i.
		invoiceRepoMysql.
		UpdateInvoiceCategory(
			ctx,
			user.ID,
			receiptCode,
			2); err != nil {
		return err
	}

	return nil
}

func (i *InvoiceUseCase) GetCompletedInvoiceProductsByReceiptCode(ctx context.Context,
	email string,
	receiptCode string) ([]domain.CompletedInvoiceProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	res, err := i.invoiceRepoMysql.GetCompletedInvoiceProductByReceiptNumber(
		ctx,
		user.ID,
		receiptCode)

	return res, nil
}
