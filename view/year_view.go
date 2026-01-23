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
	fmt.Println("2. Check day")
	fmt.Println("3. Check week")
	fmt.Println("4. Check month")
	fmt.Println("0. Exit")
}