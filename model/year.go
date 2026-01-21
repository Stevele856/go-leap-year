package model

import (
	"fmt"
)

func ValidateYear(year int) error {
	if year <= 0 {
		return fmt.Errorf("year must greater than 0")
	}
	return nil
}

func IsLeapYear(year int) bool {
	switch {
	case year % 400 == 0:
		return true
	case year % 100 == 0:
		return false
	case year % 4 == 0:
		return true
	default:
		return false
	}
}