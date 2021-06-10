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

type InvoiceResponse struct {
	StartDate   string `json:"start_date"`
	FinishDate  string `json:"finish_date"`
	ReceiptCode string `json:"receipt_code"`
}

type InvoiceProductResponse struct {
	ID           int    `json:"invoice_product_id"`
	Vendor       string `json:"product_vendor"`
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
	Quantity     uint   `json:"product_quantity"`
}

type InvoiceRepository interface {
	CreateCheckOut(ctx context.Context,
		startDate, finishDate time.Time,
		userID uint,
		cart []Cart) error
	UpdateInvoiceCategory(ctx context.Context,
		userID uint,
		receiptCode string,
		invoiceCategory int) error
	GetInvoiceByCategory(ctx context.Context,
		userID uint,
		invoiceCategory int) ([]InvoiceResponse, error)
	GetInvoiceProductByReceiptNumber(ctx context.Context,
		userID uint,
		receiptNumber string) ([]InvoiceProductResponse, error)
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
	GetInvoicesAccepted(ctx context.Context,
		email string) ([]InvoiceResponse, error)
	GetInvoicesOnGoing(ctx context.Context,
		email string) ([]InvoiceResponse, error)
	GetInvoicesCompleted(ctx context.Context,
		email string) ([]InvoiceResponse, error)
	GetInvoicesByReceiptCode(ctx context.Context,
		email, receiptCode string) ([]InvoiceProductResponse, error)
}
