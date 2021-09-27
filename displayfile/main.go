package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	if len(a) == 1 {
		missing := "File name missing"
		for i := 0; i < len(missing); i++ {
			z01.PrintRune(rune(missing[i]))
		}
	} else {
		err := os.Open
		if err != nil {
			g := "Too many arguments!!"
			for i := 0; i < len(g); i++ {
				z01.PrintRune(rune(g[i]))
			}
		} else {
			w := "Almost there!!"
			for i := 0; i < len(w); i++ {
				z01.PrintRune(rune(w[i]))
			}

		}

	}
}
