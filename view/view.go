package view

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/check-leap-year/types"
	"github.com/check-leap-year/validator"
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

func ShowProgram(scanner *bufio.Scanner, prompt string) (int, error) {
	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			ShowError(fmt.Errorf("cannot read input"))
		}
		input := strings.TrimSpace(scanner.Text())
		menu, err := strconv.Atoi(input)

		if err != nil {
			ShowError(fmt.Errorf("menu must be in integer"))
			continue
		}

		if err := validator.ValidateMenu(menu); err != nil {
			ShowError(err)
			continue
		}

		return menu, nil
	}

}

func ShowDay(d *types.Date, weekday int) {
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Printf("%d/%d/%d is %s\n", d.Day, d.Month, d.Year, weekdays[weekday])
}

func ListPrograms() {
	fmt.Println()
	fmt.Println("========= Program Lists ===========")
	fmt.Println()
	fmt.Println("[1] Check leap year")
	fmt.Println("[2] Check day in week")
	fmt.Println("[3] Check week in year")
	fmt.Println("[0] Exit program")
	fmt.Println()
}


