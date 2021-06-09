package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
	"time"
)

type InvoiceRepoMysql struct {
	DB *gorm.DB
}

func NewInvoiceRepoMysql(db *gorm.DB) domain.InvoiceRepository {
	return &InvoiceRepoMysql{DB: db}
}

func (i *InvoiceRepoMysql) CreateCheckOut(ctx context.Context,
	startDate, finishDate time.Time,
	userID uint,
	cart []domain.Cart) error {
	createInvoice := domain.Invoice{
		ReceiptNumber:     shortuuid.New(),
		StartDate:         startDate,
		FinishDate:        finishDate,
		UserID:            userID,
		InvoiceCategoryID: 1,
	}
	if err := i.DB.
		WithContext(ctx).
		Create(&createInvoice).Error; err != nil {
		return err
	}

	for p := 0; p < len(cart); p++ {
		invoiceProduct := domain.InvoiceProduct{
			ProductID: cart[p].ProductID,
			Quantity:  cart[p].Quantity,
			InvoiceID: createInvoice.ID,
		}
		if err := i.DB.
			WithContext(ctx).
			Create(&invoiceProduct).Error; err != nil {
			return err
		}
	}

	return nil
}
