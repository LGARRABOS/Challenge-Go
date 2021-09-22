package piscine

func AppendRange(min, max int) []int {
	w := max - min
	g := make([]int, min)
	if min >= max {
		i := 0
		g[i] = 0
	} else {
		for i := min; i < w; i++ {
			g[i] = i + 1
		}
	}
	return g
}
