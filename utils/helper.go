package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/check-leap-year/view"
)

func ValidateMenu(input int) error {
	if input < 0 || input > 4 {
		return fmt.Errorf("invalid menu, please choose again!")
	}
	return nil
}

func ReadMenu(scanner *bufio.Scanner, prompt string) (int, error) {
	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			view.ShowError(fmt.Errorf("cannot read input"))
		}
		input := strings.TrimSpace(scanner.Text())
		menu, err := strconv.Atoi(input)

		if err != nil {
			view.ShowError(fmt.Errorf("menu must be in integer"))
			continue
		}

		if err := ValidateMenu(menu); err != nil {
			view.ShowError(err)
			continue
		}

		return menu, nil
	}

}
