package piscine

func AppendRange(min, max int) []int {
	var v []int
	if min >= max {
	} else {
		for i := min; i < max; i++ {
			v = append(v, i)
		}
	}
	return v
}
