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

type InvoiceRepository interface {
	Fetch(ctx context.Context)
	FetchByID(ctx context.Context)
	CreateReview(ctx context.Context, review string)
}

type InvoiceUseCase interface {
	Fetch(ctx context.Context)
	FetchByID(ctx context.Context)
	CreateReview(ctx context.Context, review string)
}
