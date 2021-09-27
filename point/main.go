package main

import "github.com/01-edu/z01"

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

type point struct {
	x string
	y string
}

func main() {
	points := point{}

	setPoint(&points)
	z01.PrintRune(rune('x'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	for i := 0; i < len(points.x); i++ {
		z01.PrintRune(rune(points.x[i]))
	}
	z01.PrintRune(rune(','))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('y'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	for i := 0; i < len(points.y); i++ {
		z01.PrintRune(rune(points.y[i]))
	}

	//	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
