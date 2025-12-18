package models

import (
	"time"

	"github.com/google/uuid"
)

type RewardEvent struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID     string    `gorm:"uniqueIndex"`
	UserID      string    `gorm:"index"`
	StockSymbol string
	Quantity    float64
	RewardedAt  time.Time
	CreatedAt   time.Time
}
