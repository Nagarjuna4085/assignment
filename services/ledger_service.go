package services

import (
	"stocky/config"
	"stocky/models"

	"github.com/google/uuid"
)

func CreateLedgerEntries(reward models.RewardEvent) {
	stockEntry := models.LedgerEntry{
		ID:          uuid.New(),
		ReferenceID: reward.ID,
		EntryType:   "STOCK",
		StockSymbol: reward.StockSymbol,
		Amount:      reward.Quantity,
		Direction:   "CREDIT",
	}

	cashEntry := models.LedgerEntry{
		ID:          uuid.New(),
		ReferenceID: reward.ID,
		EntryType:   "CASH",
		INRAmount:   reward.Quantity * 2500,
		Direction:   "DEBIT",
	}

	feeEntry := models.LedgerEntry{
		ID:          uuid.New(),
		ReferenceID: reward.ID,
		EntryType:   "FEE",
		INRAmount:   50,
		Direction:   "DEBIT",
	}

	config.DB.Create(&stockEntry)
	config.DB.Create(&cashEntry)
	config.DB.Create(&feeEntry)
}
