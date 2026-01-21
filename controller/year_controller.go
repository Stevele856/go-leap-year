package controller

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/check-leap-year/model"
	"github.com/check-leap-year/view"
)

type YearController struct{}

func NewYearController() *YearController{
	return &YearController{}
}

// func (c *YearController) CheckLeapYear(input string){
// 	year, err := strconv.Atoi(input)
// 	if err != nil {
// 		fmt.Println("Year must be a number")
// 		return
// 	}

// 	if err := model.ValidateYear(year); err != nil{
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	if model.IsLeapYear(year){
// 		fmt.Printf("%d is a leap year \n", year)
// 	} else {
// 		fmt.Printf("%d is a not a leap year \n", year)
// 	}
// }


func (c *YearController) Run(scanner *bufio.Scanner) {
	fmt.Print("Enter a year: ")

	if !scanner.Scan() {
		view.ShowError(fmt.Errorf("cannot read input"))
		return
	}

	input := strings.TrimSpace(scanner.Text())

	year, err := strconv.Atoi(input)
	if err != nil {
		view.ShowError(fmt.Errorf("year must be a number"))
		return
	}

	if err := model.ValidateYear(year); err != nil {
		view.ShowError(err)
		return
	}

	if model.IsLeapYear(year) {
		view.ShowLeapYear(year)
	} else {
		view.ShowNotLeapYear(year)
	}
}
