package main

import (
	"os"

	"stocky/config"
	"stocky/models"
	"stocky/routes"
	"stocky/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	_ = godotenv.Load()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.RewardEvent{},
		&models.LedgerEntry{},
		&models.StockPrice{},
	)

	services.StartPriceUpdater()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
