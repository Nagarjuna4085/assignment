package routes

import (
	"stocky/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/reward", controllers.CreateReward)
	r.GET("/portfolio/:userId", controllers.GetPortfolio)
}
