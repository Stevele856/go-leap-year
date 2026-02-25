package validator

import (
	"errors"
	"fmt"

	"github.com/check-leap-year/types"
)

var (
	ErrInvalidMonth = errors.New("month must be between 1 and 12")
	ErrInvalidDay   = errors.New("day must be greater than 0")
	ErrInvalidYear  = errors.New("year must be greater than 0")
	ErrInvalidDate  = errors.New("invalid date for the given month and year")
)

func ValidateMenu(input int) error {
	if input < 0 || input > 4 {
		return fmt.Errorf("invalid menu, please choose again!")
	}
	return nil
}

func ValidateDate(d *types.Date) error {
	if d.Year <= 0 {
		return ErrInvalidYear
	}

	if d.Month < 1 || d.Month > 12 {
		return ErrInvalidMonth
	}

	if d.Day < 1 {
		return ErrInvalidDate
	}

	maxDay := d.MaxDaysInMonth()
	if d.Day > maxDay {
		return fmt.Errorf("%w: expected 1-%d, got %d", ErrInvalidDate, maxDay, d.Day)
	}

	return nil
}