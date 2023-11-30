package internal

import (
	"errors"
	"strconv"
	"strings"
)

// CommandParser parses the command
func CommandParser(command string) (float64, float64, error) {
	commandSplit := strings.Split(command, " ")

	shift := 0

	volume := 0.0

	volume, err := strconv.ParseFloat(commandSplit[shift], 64)
	if err != nil {
		shift++
	}

	volume, err = strconv.ParseFloat(commandSplit[shift], 64)
	if err != nil {
		return 0, 0, err
	}

	if len(commandSplit) <= shift+1 {
		return 0, 0, errors.New("strconv.ParseFloat: parsing \"\": invalid syntax")
	}

	value, err := strconv.ParseFloat(commandSplit[shift+1], 64)

	if err != nil {
		return 0, 0, err
	}

	return volume, value, nil
}
