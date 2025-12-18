package routes

import (
	"stocky/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/reward", controllers.CreateReward)
	r.GET("/portfolio/:userId", controllers.GetPortfolio)
	r.GET("/today-stocks/:userId", controllers.GetTodayStocks)
	r.GET("/historical-inr/:userId", controllers.GetHistoricalINR)
	r.GET("/stats/:userId", controllers.GetStats)
}
