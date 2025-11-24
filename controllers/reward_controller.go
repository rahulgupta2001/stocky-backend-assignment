package controllers

import (
	"net/http"
	"time"

	"stocky-backend/models"
	"stocky-backend/services"
	"stocky-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RewardUser handles rewarding a user with stocks and creating a ledger entry
func RewardUser(c *gin.Context) {
	var req struct {
		UserID      uint    `json:"user_id"`
		StockSymbol string  `json:"stock_symbol"`
		Quantity    float64 `json:"quantity"`
		UniqueEvent string  `json:"unique_event_id"` // optional
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique event ID if not provided
	if req.UniqueEvent == "" {
		req.UniqueEvent = uuid.New().String()
	}

	// Check for duplicate reward
	var existing models.Reward
	if err := utils.DB.Where("unique_event_id = ?", req.UniqueEvent).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Reward event already exists"})
		return
	}

	// Get stock price (dummy or real service)
	price := services.GetStockPrice(req.StockSymbol)

	// Create reward record
	reward := models.Reward{
		UserID:        req.UserID,
		StockSymbol:   req.StockSymbol,
		Quantity:      req.Quantity,
		PricePerUnit:  price,
		RewardAt:      time.Now(),
		UniqueEventID: req.UniqueEvent,
	}

	if err := utils.DB.Create(&reward).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add reward"})
		return
	}

	// Calculate internal fees
	brokerage := 0.02 * price * req.Quantity // 2% brokerage
	stt := 0.001 * price * req.Quantity      // 0.1% STT
	gst := 0.18 * brokerage                  // 18% GST on brokerage
	totalFee := brokerage + stt + gst
	cashOutflow := price*req.Quantity + totalFee

	// Create ledger entry
	ledger := models.Ledger{
		RewardID:     reward.ID,
		StockSymbol:  req.StockSymbol,
		Quantity:     req.Quantity,
		PricePerUnit: price,
		CashOutflow:  cashOutflow,
		BrokerageFee: brokerage,
		STTFee:       stt,
		GSTFee:       gst,
		TotalFee:     totalFee,
		CreatedAt:    time.Now(),
	}

	if err := utils.DB.Create(&ledger).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add ledger entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Reward added successfully",
		"unique_event_id": req.UniqueEvent,
	})
}
