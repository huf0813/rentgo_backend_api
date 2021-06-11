package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
	"log"
	"time"
)

type InvoiceRepoMysql struct {
	DB *gorm.DB
}

func NewInvoiceRepoMysql(db *gorm.DB) domain.InvoiceRepository {
	return &InvoiceRepoMysql{DB: db}
}

func (i *InvoiceRepoMysql) CreateReviewByInvoiceProductID(ctx context.Context,
	invoiceProductID, star uint,
	review string) error {
	updateReview := domain.InvoiceProduct{
		Rating: star,
		Review: review,
	}
	if err := i.DB.
		WithContext(ctx).
		Where("invoice_products.id = ?", invoiceProductID).
		Updates(&updateReview).
		Error; err != nil {
		return err
	}
	return nil
}

func (i *InvoiceRepoMysql) GetInvoiceByCategory(ctx context.Context,
	userID uint,
	invoiceCategory int) ([]domain.InvoiceResponse, error) {
	var res []domain.InvoiceResponse

	query := fmt.Sprintf("SELECT i.receipt_code as receipt_code, i.start_date as start_date, i.finish_date as finish_date FROM invoices i JOIN invoice_products ip ON ip.invoice_id = i.id WHERE i.user_id = %d AND i.invoice_category_id = %d", userID, invoiceCategory)
	rows, err :=
		i.DB.WithContext(ctx).
			Raw(query).
			Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var row domain.InvoiceResponse
		if err := rows.Scan(&row.ReceiptCode,
			&row.StartDate,
			&row.FinishDate); err != nil {
			return nil, err
		}
		res = append(res, row)
	}

	return res, nil
}

func (i *InvoiceRepoMysql) GetInvoiceProductByReceiptNumber(ctx context.Context,
	userID uint,
	receiptNumber string) ([]domain.InvoiceProductResponse, error) {
	var res []domain.InvoiceProductResponse

	rows, err := i.DB.WithContext(ctx).
		Raw("select p.name as product_name, "+
			"vendor.name as product_vendor, "+
			"ip.quantity as product_quantity, "+
			"ip.id as invoice_product_id, "+
			"p.price as product_price from products p "+
			"JOIN invoice_products ip on ip.product_id = p.id "+
			"JOIN invoices i on i.id = ip.invoice_id "+
			"JOIN users user_invoice on user_invoice.id = i.user_id "+
			"JOIN users vendor on vendor.id = p.user_id "+
			"where i.receipt_code = ? AND user_invoice.id = ?",
			receiptNumber,
			userID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var row domain.InvoiceProductResponse
		if err := rows.Scan(&row.ProductName,
			&row.Vendor,
			&row.Quantity,
			&row.ID,
			&row.ProductPrice); err != nil {
			return nil, err
		}
		res = append(res, row)
	}

	return res, nil
}

func (i *InvoiceRepoMysql) GetCompletedInvoiceProductByReceiptNumber(ctx context.Context,
	userID uint,
	receiptCode string) ([]domain.CompletedInvoiceProductResponse,
	error) {
	var res []domain.CompletedInvoiceProductResponse

	rows, err := i.DB.WithContext(ctx).
		Raw("select p.name as product_name, "+
			"vendor.name as product_vendor, "+
			"ip.quantity as product_quantity, "+
			"ip.id as invoice_product_id, "+
			"p.price as product_price from products p "+
			"JOIN invoice_products ip on ip.product_id = p.id "+
			"JOIN invoices i on i.id = ip.invoice_id "+
			"JOIN users user_invoice on user_invoice.id = i.user_id "+
			"JOIN users vendor on vendor.id = p.user_id "+
			"where i.receipt_code = ? "+
			"AND user_invoice.id = ? "+
			"AND i.invoice_category_id = ?",
			receiptCode,
			userID,
			2).
		Rows()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var row domain.CompletedInvoiceProductResponse
		if err := rows.Scan(&row.ProductName,
			&row.Vendor,
			&row.Quantity,
			&row.ID,
			&row.ProductPrice); err != nil {
			return nil, err
		}

		row.IsReviewed, err = i.IsCompletedInvoiceProductByReceiptNumberReviewed(
			ctx,
			userID,
			row.ID,
			receiptCode)

		res = append(res, row)
	}

	return res, nil
}

func (i *InvoiceRepoMysql) IsCompletedInvoiceProductByReceiptNumberReviewed(ctx context.Context,
	userID uint,
	invoiceProductID int,
	receiptCode string) (bool,
	error) {
	var count int64
	if err := i.DB.
		WithContext(ctx).
		Model(&domain.InvoiceProduct{}).
		Joins("JOIN invoices ON invoices.id = invoice_products.invoice_id").
		Where("invoice_products.id = ?", invoiceProductID).
		Where("invoices.receipt_code = ?", receiptCode).
		Where("invoices.user_id = ?", userID).
		Where("invoice_products.rating IS NOT NULL").
		Where("invoice_products.review IS NOT NULL").
		Count(&count).Error; err != nil {
		return false, err
	}
	if count <= 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (i *InvoiceRepoMysql) CreateCheckOut(ctx context.Context,
	startDate, finishDate time.Time,
	userID uint,
	cart []domain.Cart) error {
	createInvoice := domain.Invoice{
		ReceiptCode:       shortuuid.New(),
		StartDate:         startDate,
		FinishDate:        finishDate,
		UserID:            userID,
		InvoiceCategoryID: 3,
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

func (i *InvoiceRepoMysql) UpdateInvoiceCategory(ctx context.Context,
	userID uint,
	receiptCode string,
	invoiceCategory int) error {
	if err := i.DB.
		Model(&domain.Invoice{}).
		WithContext(ctx).
		Joins("JOIN users ON invoices.user_id = users.id").
		Where("invoices.user_id = ?", userID).
		Where("invoices.receipt_code = ?", receiptCode).
		Update("invoices.invoice_category_id", invoiceCategory).
		Error; err != nil {
		return err
	}

	return nil
}
