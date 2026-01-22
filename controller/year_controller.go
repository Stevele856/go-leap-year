package controller

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/check-leap-year/model"
	"github.com/check-leap-year/view"
)

// type YearController struct{}

// func NewYearController() *YearController {
// 	return &YearController{}
// }

// Chịu trách nhiệm điều phối

func CheckLeapYear(scanner *bufio.Scanner) {
	for {
		view.PromptYear()

		if !scanner.Scan() {
			view.ShowError(fmt.Errorf("cannot read input"))
			return
		}

		input := strings.TrimSpace(scanner.Text())

		year, err := strconv.Atoi(input)
		if err != nil {
			view.ShowError(fmt.Errorf("year must be a number"))
			continue
		}

		if err := model.ValidateYear(year); err != nil {
			view.ShowError(err)
			continue
		}

		if model.IsLeapYear(year) {
			view.ShowLeapYear(year)
		} else {
			view.ShowNotLeapYear(year)
			continue
		}
		return
	}
}
