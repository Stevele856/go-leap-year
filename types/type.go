package types


type Date struct {
	Day   int
	Month int
	Year  int
}


// Năm nhuận - IsLeapYear() là logic liên quan đến Date → nên ở types
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func (d *Date) LeapYear() bool {
	return IsLeapYear(d.Year)
}


// Xác định số ngày tối đa của mỗi tháng (Kể cả tháng có năm nhuận/không nhuân)
func (d *Date) MaxDaysInMonth() int {
	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Nếu tháng 2 có 29 ngày => Năm nhuận
	if d.Month == 2 && d.LeapYear() {
		return 29
	}
	return daysInMonth[d.Month-1]
}
