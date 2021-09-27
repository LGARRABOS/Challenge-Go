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
			w := "Almost there!!"
			fmt.Println(w)

		} else {
			g := "Too many arguments!!"
			fmt.Println(g)
		}

	}
}
