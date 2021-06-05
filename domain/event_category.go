package domain

import (
	"gorm.io/gorm"
)

type EventCategory struct {
	gorm.Model
	Name   string  `gorm:"not null" json:"name"`
	Events []Event `gorm:"foreignKey:EventCategoryID"`
}

type EventCategoryRepository interface {
}

type EventCategoryUseCase interface {
}
