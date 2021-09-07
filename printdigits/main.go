package main

import (
	"github.com/01-edu/z01"
)

func main() {
	alphabet := "0123456789"
	for i := 0; i < len(alphabet); i++ {
		z01.PrintRune(rune(alphabet[i]))
	}

	z01.PrintRune('\n')
	}