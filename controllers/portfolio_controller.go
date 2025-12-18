package controllers

import (
	"stocky/config"
	"stocky/models"

	"github.com/gin-gonic/gin"
)

func GetPortfolio(c *gin.Context) {
	userId := c.Param("userId")

	var rewards []models.RewardEvent
	config.DB.Where("user_id = ?", userId).Find(&rewards)

	portfolio := make(map[string]float64)
	for _, r := range rewards {
		portfolio[r.StockSymbol] += r.Quantity
	}

	c.JSON(200, gin.H{"portfolio": portfolio})
}
