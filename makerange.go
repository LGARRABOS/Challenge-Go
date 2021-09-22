package piscine

func MakeRange(min, max int) []int {
	g := max - min
	v := make([]int, g)
	if max <= min {
		v = make([]int, 0)
		v = nil
	} else {
		for i := 0; i < g; i++ {
			v[i] = i + min
		}
	}
	return v
}
