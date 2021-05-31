package domain

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	Quantity   uint      `json:"quantity"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}
