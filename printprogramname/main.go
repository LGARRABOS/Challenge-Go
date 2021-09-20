package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[0]
	z := []rune(a)
	for i := 0; i < len(a); i++ {
		z01.PrintRune(z[i])
	}
}
