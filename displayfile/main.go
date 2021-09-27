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
		file, _ := os.Open("quest8.txt")
		if len(os.Args) > 2 {
			g := "Too many arguments"
			fmt.Println(g)

		} else {
			w := make([]byte, 14)
			file.Read(w)
			fmt.Println(string(w))
			file.Close()
		}
	}
}
