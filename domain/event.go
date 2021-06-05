package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name            string `gorm:"not null" json:"name"`
	PeopleAmount    string `gorm:"not null" json:"people_amount"`
	LargeArea       string `gorm:"not null" json:"large_area"`
	EventCategoryID uint   `json:"event_category_id"`
}

type EventRepository interface {
}

type EventUseCase interface {
}
