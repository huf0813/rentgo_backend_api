package domain

import (
	"gorm.io/gorm"
)

type EventProduct struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Quantity  uint   `json:"quantity"`
	ProductID uint   `json:"product_id"`
	EventID   uint   `json:"event_id"`
}

type EventProductRepository interface {
}

type EventProductUseCase interface {
}
