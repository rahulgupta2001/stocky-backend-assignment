package routes

import (
	"stocky-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api") // Use the passed engine `r`
	{
		api.POST("/reward", controllers.RewardUser)
		api.GET("/portfolio/:userId", controllers.GetPortfolio)
		api.GET("/today-stocks/:userId", controllers.GetTodayStocks)
		api.GET("/historical-inr/:userId", controllers.GetHistoricalINR)
		api.GET("/stats/:userId", controllers.GetStats)
	}
}
