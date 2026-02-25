package model

import (
	"testing"
	"time"

	"github.com/check-leap-year/types"
)

func TestCenturyAnchor(t *testing.T) {
	tests := []struct {
		name string
		year int
		want int
	}{
		{"1600", 1600, 2},
		{"2000", 2000, 2},
		{"2400", 2000, 2},

		{"1700", 1700, 0},
		{"2100", 2100, 0},

		{"1800", 1800, 5},
		{"2200", 2200, 5},

		{"1900", 1900, 3},
		{"2300", 2300, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CenturyAnchor(tt.year)
			if got != tt.want{
				t.Errorf("CenturyAnchor(%d) = %d, want %d", tt.year, got, tt.want)
			}
		})
	}
}


func TestDoomsdayOfYear(t *testing.T){
	tests := []struct{
		name string
		year int
	}{
		{"Year 2023", 2023},
		{"Year 2024 (leap)", 2024},
		{"Year 2000 (century leap)", 2000},
		{"Year 1900 (not leap century)", 1900},
		{"Year 2100 (not leap century)", 2100},

	}

	for _, tt := range tests {
		t.Run(tt.name,func(t *testing.T){
			got := DoomsdayOfYear(tt.year)

			expected := int(time.Date(
				tt.year,
				time.April, 4, 
				0, 0, 0, 0, 
				time.UTC,
			).Weekday())

			if got != expected {
				t.Errorf("DoomsdayOfYear(%d) = %d, want %d", tt.year, got, expected)
			}
		})
	}
}

func TestMonthDoomsday(t *testing.T){
	tests := []struct{
		name string
		month int
		isLeap bool
		want int
	}{
		// January
		{"Jan non-leap", 1, false, 3},
		{"Jan leap", 1, true, 4},

		// Febuary
		{"Feb non-leap", 2, false, 28},
		{"Feb leap", 2, true, 29},

		// Fixed month
		{"Mar", 3, false, 14},
		{"April", 4, false, 4},
		{"May", 5, false, 9},
		{"June", 6, false, 6},
		{"July", 7, false, 11},
		{"August", 8, false, 8},
		{"September", 9, false, 5},
		{"October", 10, false, 10},
		{"November", 11, false, 7},
		{"December", 12, false, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MonthDoomsday(tt.month, tt.isLeap)
			if got != tt.want{
				t.Errorf("MonthDoomsday(%d, %v) = %d, want %d", tt.month, tt.isLeap, got, tt.want)
			}
		})
	}
}

func TestWeekdayFromDate(t *testing.T) {
	tests := []struct {
		name string
		date types.Date
	}{
		{"2026-02-25", types.Date{Year: 2026, Month: 02, Day: 25}},
		{"2024-02-29 (Leap)", types.Date{Year: 2024, Month: 02, Day: 29}},
		{"2000-01-01 (Century leap)", types.Date{Year: 2000, Month: 01, Day: 01}},
		{"1900-01-01 (not leap century)", types.Date{Year: 1900, Month: 01, Day: 01}},
		{"2026-12-31", types.Date{Year: 2026, Month: 12, Day: 31}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WeekdayFromDate(&tt.date)

			expected := int(time.Date(
				tt.date.Year,
				time.Month(tt.date.Month),
				tt.date.Day,
				0,0,0,0,
				time.UTC,
			).Weekday())

			if got != expected {
				t.Errorf("WeeldayFromDate(%v) = %d, want %d", tt.date, got, expected)
			}
		})
	}
}