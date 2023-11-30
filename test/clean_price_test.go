package test

import (
	"Pricer/internal"
	"testing"
)

func TestCalculateCleanPrice(t *testing.T) {
	volume := 10.0
	price := 5.0

	expected := 500.0
	result := internal.CalculateCleanPrice(volume, price)

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}
