package controller

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/check-leap-year/model"
	"github.com/check-leap-year/types"
	"github.com/check-leap-year/view"
)

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

		if types.IsLeapYear(year) {
			view.ShowLeapYear(year)
		} else {
			view.ShowNotLeapYear(year)
		}
		return
	}
}

func CheckDateInWeek(scanner *bufio.Scanner) {
	for {
		view.PromptYear()

		if !scanner.Scan() {
			view.ShowError(fmt.Errorf("cannot read input"))
			return
		}

		yearStr := strings.TrimSpace(scanner.Text())
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			view.ShowError(fmt.Errorf("year must be a number"))
			continue
		}

		view.PromptMonth()

		if !scanner.Scan() {
			view.ShowError(fmt.Errorf("cannot read input"))
			return
		}

		monthStr := strings.TrimSpace(scanner.Text())
		month, err := strconv.Atoi(monthStr)
		if err != nil {
			view.ShowError(fmt.Errorf("month must be a number"))
			continue
		}

		view.PromptDay()

		if !scanner.Scan() {
			view.ShowError(fmt.Errorf("cannot read input"))
			return
		}

		dayStr := strings.TrimSpace(scanner.Text())
		day, err := strconv.Atoi(dayStr)
		if err != nil {
			view.ShowError(fmt.Errorf("day must be a number"))
			continue
		}

		date, err := model.NewDate(day, month, year)
		if err != nil {
			view.ShowError(err)
			continue
		}

		weekday := model.WeekdayFromDate(date)
		view.ShowWeekday(date, weekday)
	}
}
