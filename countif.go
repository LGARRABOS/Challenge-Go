package piscine

func CountIf(f func(string) bool, tab []string) int {
	j := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) == true {
			j = j + 1
		}
	}
	return j
}
