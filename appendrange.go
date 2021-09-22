package piscine

func AppendRange(min, max int) []int {
	w := max - min
	g := make([]int, min)
	for i := min; i < w; i++ {
		g[i] = i + 1
	}
	return g
}
