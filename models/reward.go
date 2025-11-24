package models

import "time"

type Reward struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	StockSymbol   string
	Quantity      float64
	PricePerUnit  float64
	RewardAt      time.Time
	UniqueEventID string `gorm:"uniqueIndex"`
}
