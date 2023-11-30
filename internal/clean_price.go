package internal

// CalculateCleanPrice
func CalculateCleanPrice(volume float64, price float64) float64 {
	cleanPrice := price * 1000.0 / volume
	return cleanPrice
}
