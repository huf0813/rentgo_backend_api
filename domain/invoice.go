package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Invoice struct {
	gorm.Model
	ReceiptCode       string           `gorm:"unique" json:"receipt_code"`
	StartDate         time.Time        `json:"start_date"`
	FinishDate        time.Time        `json:"finish_date"`
	UserID            uint             `json:"user_id"`
	InvoiceCategoryID uint             `json:"invoice_category_id"`
	InvoiceProducts   []InvoiceProduct `gorm:"foreignKey:InvoiceID" json:"invoice_products"`
}

type InvoiceCheckoutRequest struct {
	CartIDS    []int  `json:"cart_ids"`
	StartDate  string `json:"start_date"`
	FinishDate string `json:"finish_date"`
}

type InvoiceRepository interface {
	CreateCheckOut(ctx context.Context,
		startDate, finishDate time.Time,
		userID uint,
		cart []Cart) error
	UpdateInvoiceOnGoing(ctx context.Context,
		userID uint,
		receiptCode string) error
	UpdateInvoiceCompleted(ctx context.Context,
		userID uint,
		receiptCode string) error
}

type InvoiceUseCase interface {
	CreateCheckOut(ctx context.Context,
		startDate, finishDate time.Time,
		email string,
		cartIDS []int) error
	UpdateInvoiceOnGoing(ctx context.Context,
		email string,
		receiptCode string) error
	UpdateInvoiceCompleted(ctx context.Context,
		email string,
		receiptCode string) error
}
