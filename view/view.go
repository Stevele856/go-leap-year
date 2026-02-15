package view

import (
	"fmt"
	"github.com/check-leap-year/types"
)

func ShowLeapYear(year int) {
	fmt.Printf("%d is a leap year\n", year)
}

func ShowNotLeapYear(year int) {
	fmt.Printf("%d is not a leap year\n", year)
}

func ShowError(err error) {
	fmt.Println("Error:", err.Error())
}

func PromptYear() {
	fmt.Print("Enter a year: ")
}

func PromptMonth() {
	fmt.Print("Enter a month (1-12): ")
}

func PromptDay() {
	fmt.Print("Enter a day: ")
}

func ShowWeekday(d *types.Date, weekday int) {
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Printf("%d/%d/%d is %s\n", d.Day, d.Month, d.Year, weekdays[weekday])
}

func ListPrograms() {
	fmt.Println()
	fmt.Println("========= Program Lists ===========")
	fmt.Println()
	fmt.Println("[1] Check leap year")
	fmt.Println("[2] Check day in week")
	fmt.Println()
	// fmt.Println("3. Check week in year")
	// fmt.Println("4. Check day in year")
	// fmt.Println("5. Check quater in year")
	// fmt.Println("6. Check day of end of year")
	// fmt.Println("7. Calculate day")
	// fmt.Println("0. Exit")
}
