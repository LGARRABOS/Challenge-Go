package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args
	if len(a) == 1 {
		missing := "File name missing"
		fmt.Println(missing)
	} else {
		err := os.Open
		if err != nil {
			g := "Too many arguments!!"
			fmt.Println(g)

		} else {
			w := "Almost there!!"
			fmt.Println(w)

		}

	}
}
