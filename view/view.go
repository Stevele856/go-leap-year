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

func PromptYear() {
    fmt.Print("Enter a year: ")
}

func ShowMenu(){
	fmt.Println("1. Check leap year")
	fmt.Println("2. Check day in week")
	fmt.Println("3. Check week in year")
	fmt.Println("4. Check day in year")
	fmt.Println("5. Check quater in year")
	fmt.Println("6. Check day of end of year")
	fmt.Println("7. Caculate day")
	fmt.Println("0. Exit")
}