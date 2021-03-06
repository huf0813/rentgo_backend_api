package domain

import "gorm.io/gorm"

type InvoiceProduct struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	InvoiceID uint   `json:"invoice_id"`
	Quantity  uint   `json:"quantity"`
	Rating    uint   `gorm:"default:null" json:"rating"`
	Review    string `gorm:"default:null" json:"review"`
}
