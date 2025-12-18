package models

import (
	"time"

	"github.com/google/uuid"
)

type LedgerEntry struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	ReferenceID uuid.UUID
	EntryType   string // STOCK, CASH, FEE
	StockSymbol string
	Amount      float64 // stock units
	INRAmount   float64
	Direction   string // DEBIT / CREDIT
	CreatedAt   time.Time
}
