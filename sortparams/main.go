package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := os.Args
	for x := 0; x < len(i); x++ {
		for w := 2; w < len(i); w++ {
			if i[w] < i[w-1] {
				c := i[w]
				i[w] = i[w-1]
				i[w-1] = c
			}
		}
	}
	for a := 1; a < len(i); a++ {
		for _, j := range i[a] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
