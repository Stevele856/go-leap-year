package model

import "github.com/check-leap-year/types"

// Hàm tính mốc thế kỷ
func CenturyAnchor(year int) int {
	century := year / 100
	switch century % 4 {
	case 0:
		return 2
	case 1:
		return 0
	case 2:
		return 5
	default:
		return 3
	}
}

// Hàm tính ngày Doomsday của năm
func DoomsdayOfYear(year int) int {
	y := year % 100
	a := y / 12
	b := y % 12
	c := b / 4
	d := a + b + c

	return (d + CenturyAnchor(year)) % 7
}

// Hàm trả về ngày Doomsday của tháng
func MonthDoomsday(month int, leap bool) int {
	switch month {
	case 1:
		if leap {
			return 4
		}
		return 3
	case 2:
		if leap {
			return 29
		}
		return 28
	case 3:
		return 14
	case 4:
		return 4
	case 5:
		return 9
	case 6:
		return 6
	case 7:
		return 11
	case 8:
		return 8
	case 9:
		return 5
	case 10:
		return 10
	case 11:
		return 7
	case 12:
		return 12
	default:
		return 0
	}
}


// Hàm chính để tính ngày trong tuần từ ngày, tháng, năm
func WeekdayFromDate(d *types.Date) int {
	doomsday := DoomsdayOfYear(d.Year)
	anchor := MonthDoomsday(d.Month, types.IsLeapYear(d.Year))
	delta := (d.Day - anchor) % 7
	weekday := (doomsday + delta + 7) % 7
	return weekday
}
