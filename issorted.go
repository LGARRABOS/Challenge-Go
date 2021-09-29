package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	j := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i]+1) >= 0 {
			j = j + 1
		} else {
			break
		}
	}
	if j == len(a)-1 {
		return true
	} else {
		j = 0
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i]+1) <= 0 {
				j = j + 1
			}
		}
	}
	if j == len(a)-1 {
		return true
	}
	return false
}
