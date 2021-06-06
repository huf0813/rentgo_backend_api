package domain

import "gorm.io/gorm"

type InvoiceCategory struct {
	gorm.Model
	Name     string    `gorm:"unique;not null" json:"name"`
	Invoices []Invoice `gorm:"foreignKey:InvoiceCategoryID" json:"invoices"`
}
