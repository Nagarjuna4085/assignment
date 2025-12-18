package controllers

import (
	"math"
	"stocky/config"
	"stocky/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPortfolio(c *gin.Context) {
	userId := c.Param("userId")

	var rewards []models.RewardEvent
	config.DB.Where("user_id = ?", userId).Find(&rewards)

	portfolio := make(map[string]float64)
	totalINR := 0.0

	for _, r := range rewards {
		portfolio[r.StockSymbol] += r.Quantity

		// Fetch latest price
		var price models.StockPrice
		err := config.DB.Where("stock_symbol = ?", r.StockSymbol).Order("price_time desc").First(&price).Error
		if err == nil {
			totalINR += r.Quantity * price.Price
		}
	}
	totalINR = math.Round(totalINR*100) / 100

	c.JSON(200, gin.H{
		"user_id":   userId,
		"portfolio": portfolio,
		"total_inr": totalINR,
	})
}

// GET /today-stocks/:userId
func GetTodayStocks(c *gin.Context) {
	userId := c.Param("userId")
	start := time.Now().UTC().Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	var rewards []models.RewardEvent
	config.DB.Where("user_id = ? AND rewarded_at >= ? AND rewarded_at < ?", userId, start, end).Find(&rewards)

	c.JSON(200, gin.H{
		"user_id": userId,
		"date":    start.Format("2006-01-02"),
		"rewards": rewards,
	})
}

// GET /historical-inr/:userId
func GetHistoricalINR(c *gin.Context) {
	userId := c.Param("userId")

	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now().AddDate(0, 0, -1).UTC().Truncate(24 * time.Hour)

	var rewards []models.RewardEvent
	config.DB.Where("user_id = ? AND rewarded_at >= ? AND rewarded_at <= ?", userId, start, end).Find(&rewards)

	history := make(map[string]float64)
	for _, r := range rewards {
		var price models.StockPrice
		config.DB.Where("stock_symbol = ?", r.StockSymbol).Order("price_time desc").First(&price)

		dateKey := r.RewardedAt.Format("2006-01-02")
		history[dateKey] += math.Round(r.Quantity*price.Price*100) / 100
	}

	c.JSON(200, gin.H{
		"user_id": userId,
		"history": history,
	})
}

// GET /stats/:userId
func GetStats(c *gin.Context) {
	userId := c.Param("userId")

	// Total shares rewarded today
	start := time.Now().UTC().Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	var rewards []models.RewardEvent
	config.DB.Where("user_id = ? AND rewarded_at >= ? AND rewarded_at < ?", userId, start, end).Find(&rewards)

	todayTotals := make(map[string]float64)
	for _, r := range rewards {
		todayTotals[r.StockSymbol] += r.Quantity
	}

	// Current portfolio value
	var allRewards []models.RewardEvent
	config.DB.Where("user_id = ?", userId).Find(&allRewards)

	totalINR := 0.0
	for _, r := range allRewards {
		var price models.StockPrice
		config.DB.Where("stock_symbol = ?", r.StockSymbol).Order("price_time desc").First(&price)
		totalINR += r.Quantity * price.Price
	}
	totalINR = math.Round(totalINR*100) / 100

	c.JSON(200, gin.H{
		"user_id":               userId,
		"today_rewards":         todayTotals,
		"current_portfolio_inr": totalINR,
	})
}
