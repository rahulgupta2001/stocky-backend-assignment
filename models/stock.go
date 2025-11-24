package models

type Stock struct {
	Symbol      string `gorm:"primaryKey" json:"symbol"`
	CompanyName string `json:"company_name"`
}
