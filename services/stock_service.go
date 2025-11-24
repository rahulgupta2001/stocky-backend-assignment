package services

import (
	"math/rand"
)

func GetStockPrice(symbol string) float64 {
	// Dummy price between 1000 to 5000
	return 1000 + rand.Float64()*(5000-1000)
}
