package model

import "fmt"

type Date struct {
	Day   int
	Month int
	Year  int
}

func (d Date) Validate() error {
	// Kiểm tra tháng hợp lệ (1-12) 1 2 3 4 5 6 7 8 9 10 11 12
	if d.Month < 1 || d.Month > 12 {
		return fmt.Errorf("month from 1-12")
	}
	// Kiểm tra ngày hợp lệ (1-31)
	if d.Day < 1 {
		return fmt.Errorf("day must greater than 0")
	}
	// Số ngày tối đa của mỗi tháng => slice
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Kiểm tra năm hợp lệ
	if d.Year <= 0 {
		return fmt.Errorf("year must greater than 0")
	}
	// Kiểm tra năm nhuận
	if isLeapYear(d.Year) {
		daysInMonth[1] = 29 // Năm nhuận thì tháng 2 có 29 ngày
	}

	// Kiểm tra ngày có vượt quá số ngày trong tháng không
	// 31 > 30 true (Mock month 11 )
	if d.Day > daysInMonth[d.Month-1] {
		return fmt.Errorf("invalid day for this month")
	}
	return nil // Hợp lệ
}

// Hàm phụ kiểm tra năm nhuận
func isLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
