package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) int {
	f := nbr % 2
	if f == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) bool {
	if even(nbr) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(len(os.Args)) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
