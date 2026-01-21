package view

import "fmt"



func ShowLeapYear(year int) {
	fmt.Printf("%d is a leap year\n", year)
}

func ShowNotLeapYear(year int) {
	fmt.Printf("%d is not a leap year\n", year)
}

func ShowError(err error) {
	fmt.Println(err.Error())
}