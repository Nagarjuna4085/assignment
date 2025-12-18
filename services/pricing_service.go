package services

import (
	"math/rand"
	"time"

	"stocky/config"
	"stocky/models"
)

func StartPriceUpdater() {
	go func() {
		for {
			symbols := []string{"RELIANCE", "TCS", "INFOSYS"}

			for _, s := range symbols {
				price := rand.Float64()*3000 + 500
				config.DB.Create(&models.StockPrice{
					StockSymbol: s,
					Price:       price,
					PriceTime:   time.Now(),
				})
			}
			time.Sleep(1 * time.Hour)
		}
	}()
}
