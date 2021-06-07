package usecase

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"time"
)

type InvoiceUseCase struct {
	invoiceRepoMysql domain.InvoiceRepository
	timeOut          time.Duration
}

func NewInvoiceUseCase(i domain.InvoiceUseCase, timeOut time.Duration) domain.InvoiceUseCase {
	return &InvoiceUseCase{
		invoiceRepoMysql: i,
		timeOut:          timeOut,
	}
}

func (i InvoiceUseCase) Fetch(ctx context.Context) {

}

func (i InvoiceUseCase) FetchByID(ctx context.Context) {
	panic("implement me")
}

func (i InvoiceUseCase) CreateReview(ctx context.Context, review string) {
	panic("implement me")
}
