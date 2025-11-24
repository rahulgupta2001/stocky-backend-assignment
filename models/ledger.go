package models

import (
	"time"

	"gorm.io/gorm"
)

type Ledger struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	RewardID     uint           `json:"reward_id"` // FK to Reward
	StockSymbol  string         `json:"stock_symbol"`
	Quantity     float64        `json:"quantity"`
	PricePerUnit float64        `json:"price_per_unit"`
	CashOutflow  float64        `json:"cash_outflow"`
	BrokerageFee float64        `json:"brokerage_fee"`
	STTFee       float64        `json:"stt_fee"`
	GSTFee       float64        `json:"gst_fee"`
	TotalFee     float64        `json:"total_fee"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
