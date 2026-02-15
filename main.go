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
	for {
		view.ListPrograms()
		menu, err := utils.ReadMenu(scanner, "Select program: ")
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch menu {
		case 1:
			controller.CheckLeapYear(scanner)
		case 2:
			controller.CheckDateInWeek(scanner)
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid program, please choose again!")
			continue
		}
		
	}
}
