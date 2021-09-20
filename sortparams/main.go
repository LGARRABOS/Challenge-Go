package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := os.Args
	for x := 0; x < 1000; x++ {
		for w := 1; w < len(i); w++ {
			if i[0] < i[w-1] {
				c := i[w]
				i[w] = i[w-1]
				i[w-1] = c
			} else {
				continue
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
