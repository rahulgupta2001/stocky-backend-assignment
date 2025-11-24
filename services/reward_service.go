package services

import (
	"errors"
	"stocky-backend/models"
	"stocky-backend/utils"
	"time"
)

func RewardUser(userID uint, stockSymbol string, quantity float64, eventID string) error {
	// Check for duplicate event
	var existing models.Reward
	if err := utils.DB.Where("unique_event_id = ?", eventID).First(&existing).Error; err == nil {
		return errors.New("duplicate reward event")
	}

	// Create reward
	reward := models.Reward{
		UserID:        userID,
		StockSymbol:   stockSymbol,
		Quantity:      quantity,
		RewardAt:      time.Now(),
		UniqueEventID: eventID,
	}
	if err := utils.DB.Create(&reward).Error; err != nil {
		return err
	}

	// Add to ledger (dummy fees calculation)
	cashOutflow := quantity * GetStockPrice(stockSymbol)
	brokerage := cashOutflow * 0.001
	stt := cashOutflow * 0.00025
	gst := (brokerage + stt) * 0.18
	totalFee := brokerage + stt + gst

	ledger := models.Ledger{
		RewardID:     reward.ID,
		StockSymbol:  stockSymbol,
		Quantity:     quantity,
		CashOutflow:  cashOutflow,
		BrokerageFee: brokerage,
		STTFee:       stt,
		GSTFee:       gst,
		TotalFee:     totalFee,
		CreatedAt:    time.Now(),
	}
	return utils.DB.Create(&ledger).Error
}
