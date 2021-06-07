package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type InvoiceRepoMysql struct {
	DB *gorm.DB
}

func NewInvoiceRepoMysql(db *gorm.DB) domain.InvoiceRepository {
	return &InvoiceRepoMysql{DB: db}
}

func (i *InvoiceRepoMysql) Fetch(ctx context.Context) {

}

func (i *InvoiceRepoMysql) FetchByID(ctx context.Context) {
	panic("implement me")
}

func (i *InvoiceRepoMysql) CreateReview(ctx context.Context, review string) {
	panic("implement me")
}
