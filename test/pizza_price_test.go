package test

import (
	"Pricer/internal"
	"testing"
)

func TestCalculatePizzaPrice(t *testing.T) {
	diameter := 23.0
	price := 635.0

	expectedPrice := 1.53
	expectedArea := 415.48

	area, piecePrice, err := internal.CalculatePizzaPrice(diameter, price)

	if piecePrice != expectedPrice {
		t.Errorf("Expected %.2f, but got %.2f", expectedPrice, piecePrice)
	}
	if area != expectedArea {
		t.Errorf("Expected %.2f, but got %.2f", expectedArea, area)
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
