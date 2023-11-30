package internal

import "math"

// Расчет стоимость пиццы
func CalculatePizzaPrice(diameter float64, price float64) (float64, float64, error) {
	area := math.Pi * math.Pow(diameter/2.0, 2.0)
	cleanPrice := price / area
	return math.Round(area*100) / 100, math.Round(cleanPrice*100) / 100, nil

}
