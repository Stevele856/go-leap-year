package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/check-leap-year/controller"
	"github.com/check-leap-year/utils"
	"github.com/check-leap-year/view"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	view.ShowMenu()
	menu, err := utils.ReadMenu(scanner, "Enter menu: ")
	if err != nil {
		fmt.Println(err)
	}
	switch menu {
	case 1:
		controller.CheckLeapYear(scanner)
	}

}
