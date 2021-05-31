package domain

import (
	"gorm.io/gorm"
	"time"
)

type Invoice struct {
	gorm.Model
	Quantity   uint      `json:"quantity"`
	StartDate  time.Time `json:"start_date"`
	FinishDate time.Time `json:"finish_date"`
	UserID     uint      `json:"user_id"`
}
