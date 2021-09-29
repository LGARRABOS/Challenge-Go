package piscine

func Map(f func(int) bool, a []int) []bool {
	j := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			j[i] = true
		} else {
			j[i] = false
		}
	}
	return j
}
