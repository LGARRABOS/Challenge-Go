package piscine

func AppendRange(min, max int) []int {
	w := max - min
	g := 0
	if min >= max {
		g = 0
	} else {
		for i := min; i < w-1; i++ {
			g = i + 1
		}
	}
	return g
}
