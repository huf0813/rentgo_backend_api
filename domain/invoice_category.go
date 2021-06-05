package domain

import "gorm.io/gorm"

// cart, to_pay, on_going, completed
type InvoiceCategory struct {
	gorm.Model
	Name     string    `json:"name"`
	Invoices []Invoice `gorm:"foreignKey:InvoiceCategoryID" json:"invoices"`
}
