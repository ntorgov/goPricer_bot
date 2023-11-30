package internal

import (
	"errors"
	"math"
)

// Расчет стоимость пиццы
func CalculatePizzaPrice(diameter float64, price float64) (float64, float64, error) {
	if diameter <= 0.0 || price <= 0.0 {
		return 0, 0, errors.New("Диаметр и цена должны быть больше нуля")
	}
	area := math.Pi * math.Pow(diameter/2.0, 2.0)
	cleanPrice := price / area
	return math.Round(area*100) / 100, math.Round(cleanPrice*100) / 100, nil

}
