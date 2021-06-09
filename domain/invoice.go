package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Invoice struct {
	gorm.Model
	ReceiptNumber     string           `json:"receipt_number"`
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
}

type InvoiceUseCase interface {
	CreateCheckOut(ctx context.Context,
		startDate, finishDate time.Time,
		email string,
		cartIDS []int) error
}
