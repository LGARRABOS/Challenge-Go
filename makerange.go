package piscine

func MakeRange(min, max int) []int {
	g := max - min
	if max <= min {
		v := make([]int, 0)
		v = nil
		return v
	} else {
		v := make([]int, g)
		for i := 0; i < g; i++ {
			v[i] = i + min
		}
		return v
	}
}
