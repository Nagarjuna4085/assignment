package models

import "time"

type StockPrice struct {
	StockSymbol string
	Price       float64
	PriceTime   time.Time
}
