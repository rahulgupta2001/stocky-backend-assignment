package controllers

import (
	"net/http"
	"stocky-backend/models"
	"stocky-backend/services"
	"stocky-backend/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTodayStocks returns all stocks rewarded today for a user
func GetTodayStocks(c *gin.Context) {
	// Trim any whitespace/newlines from the URL parameter
	userId := strings.TrimSpace(c.Param("userId"))

	var rewards []models.Reward
	today := time.Now().Format("2006-01-02")

	if err := utils.DB.Where("user_id = ? AND DATE(reward_at) = ?", userId, today).Find(&rewards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch today's stocks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":      userId,
		"today_stocks": rewards,
	})
}

// GetHistoricalINR returns past INR values of a user's rewards (up to yesterday)
func GetHistoricalINR(c *gin.Context) {
	userId := strings.TrimSpace(c.Param("userId"))

	var rewards []models.Reward
	if err := utils.DB.Where("user_id = ? AND DATE(reward_at) < CURRENT_DATE", userId).Find(&rewards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch historical rewards"})
		return
	}

	historical := make(map[string]float64)
	for _, r := range rewards {
		date := r.RewardAt.Format("2006-01-02")
		price := services.GetStockPrice(r.StockSymbol)
		historical[date] += r.Quantity * price
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userId, "historical_inr": historical})
}

// GetStats returns today's reward summary and current portfolio value
func GetStats(c *gin.Context) {
	userId := strings.TrimSpace(c.Param("userId"))

	var rewards []models.Reward
	if err := utils.DB.Where("user_id = ?", userId).Find(&rewards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rewards"})
		return
	}

	today := time.Now().Format("2006-01-02")
	todayStats := make(map[string]float64)
	portfolioValue := 0.0

	for _, r := range rewards {
		price := services.GetStockPrice(r.StockSymbol)
		portfolioValue += r.Quantity * price

		if r.RewardAt.Format("2006-01-02") == today {
			todayStats[r.StockSymbol] += r.Quantity
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":       userId,
		"today_rewards": todayStats,
		"portfolio_inr": portfolioValue,
	})
}

// GetPortfolio returns current portfolio value for a user
func GetPortfolio(c *gin.Context) {
	userId := strings.TrimSpace(c.Param("userId"))

	var rewards []models.Reward
	if err := utils.DB.Where("user_id = ?", userId).Find(&rewards).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio := make(map[string]float64)
	for _, r := range rewards {
		price := services.GetStockPrice(r.StockSymbol)
		portfolio[r.StockSymbol] += r.Quantity * price
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":   userId,
		"portfolio": portfolio,
		"timestamp": time.Now(),
	})
}
