package main

import (
	"bufio"
	"os"

	"github.com/check-leap-year/controller"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	controller.CheckLeapYear(scanner)
}
