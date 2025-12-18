package controllers

import (
	"time"

	"stocky/config"
	"stocky/models"
	"stocky/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RewardRequest struct {
	EventID     string  `json:"event_id"`
	UserID      string  `json:"user_id"`
	StockSymbol string  `json:"stock_symbol"`
	Quantity    float64 `json:"quantity"`
	Timestamp   string  `json:"timestamp"`
}

func CreateReward(c *gin.Context) {
	var req RewardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existing models.RewardEvent
	if err := config.DB.Where("event_id = ?", req.EventID).First(&existing).Error; err == nil {
		c.JSON(200, gin.H{"message": "duplicate event ignored"})
		return
	}

	t, _ := time.Parse(time.RFC3339, req.Timestamp)

	reward := models.RewardEvent{
		ID:          uuid.New(),
		EventID:     req.EventID,
		UserID:      req.UserID,
		StockSymbol: req.StockSymbol,
		Quantity:    req.Quantity,
		RewardedAt:  t,
	}

	config.DB.Create(&reward)
	services.CreateLedgerEntries(reward)

	logrus.Info("Reward created")

	c.JSON(200, gin.H{"status": "success"})
}
