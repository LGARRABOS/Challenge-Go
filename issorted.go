package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	j := 0
	for i := 0; i < len(a); i++ {
		if i < i+1 {
			j = j + 1
		} else {
			break
		}
	}
	if j == len(a) {
		return true
	} else {
		return false
	}
}
