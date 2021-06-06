package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name            string         `gorm:"not null" json:"name"`
	PeopleAmount    string         `gorm:"not null" json:"people_amount"`
	LargeArea       string         `gorm:"not null" json:"large_area"`
	UserID          uint           `json:"user_id"`
	EventCategoryID uint           `json:"event_category_id"`
	EventProducts   []EventProduct `gorm:"foreignKey:EventID" json:"event_products"`
}

type EventRepository interface {
}

type EventUseCase interface {
}
