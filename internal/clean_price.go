package internal

/**
 * CalculateCleanPrice calculates the clean price
 */
func CalculateCleanPrice(volume float64, price float64) float64 {
	cleanPrice := volume * 1000 / price
	return cleanPrice
}
