package model

import (
	"github.com/check-leap-year/types"
	"github.com/check-leap-year/validator"
)

func NewDate(day, month, year int) (*types.Date, error) {
	d := &types.Date{
		Day:   day,
		Month: month,
		Year:  year,
	}
	if err := validator.ValidateDate(d); err != nil {
		return nil, err
	}
	return d, nil
}


