package test

import (
	"Pricer/internal"
	"testing"
)

func TestParserWithCleanValues(t *testing.T) {
	command := "852 150"

	volume, value, err := internal.CommandParser(command)
	expectedVolume := 852.0
	expectedValue := 150.0

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if volume != expectedVolume {
		t.Errorf("Expected %f, but got %f", expectedVolume, volume)
	}

	if value != expectedValue {
		t.Errorf("Expected %f, but got %f", expectedVolume, volume)
	}
}

func TestParserWithCommandValues(t *testing.T) {
	command := "цена 852 150"

	volume, value, err := internal.CommandParser(command)
	expectedVolume := 852.0
	expectedValue := 150.0

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if volume != expectedVolume {
		t.Errorf("Expected %f, but got %f", expectedVolume, volume)
	}

	if value != expectedValue {
		t.Errorf("Expected %f, but got %f", expectedVolume, volume)
	}
}

// Test parsing error string
func TestCapturingErrorsWithCommand(t *testing.T) {
	command := "цена 8пять2 150"

	_, _, err := internal.CommandParser(command)

	if err == nil && err.Error() != "strconv.ParseFloat: parsing \"8пять2\": invalid syntax" {
		t.Errorf("Error: %s", err)
	}

	command = "цена 852 1пять0"

	if err == nil && err.Error() != "strconv.ParseFloat: parsing \"1пять0\": invalid syntax" {
		t.Errorf("Error: %s", err)
	}
}

// Test parsing error string
func TestCapturingErrorsWithoutCommand(t *testing.T) {
	command := "8пять2 150"

	_, _, err := internal.CommandParser(command)

	if err == nil && err.Error() != "strconv.ParseFloat: parsing \"8пять2\": invalid syntax" {
		t.Errorf("Error: %s", err)
	}

	command = "852 1пять0"

	_, _, err = internal.CommandParser(command)

	if err == nil && err.Error() != "strconv.ParseFloat: parsing \"1пять0\": invalid syntax" {
		t.Errorf("Error: %s", err)
	}
}
