package piscine

func MakeRange(min, max int) []int {
	g := max - min
	v := make([]int, min)
	if max <= min {
		v = make([]int, 0)
	} else {
		for i := 0; i < g; i++ {
			v[i] = i + min
		}
	}
	return v
}
