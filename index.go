package piscine

func Index(s string, toFind string) int {
	w := 0
	z := []rune(toFind)
	l := []rune(s)
	for q, a := range s {
		if w == 0 {
			if a == z[0] {
				w = q
				break
			} else if q == len(s)-1 {
				return -1
			}
		}
	}
	for i := 0; i < len(z); i++ {
		if l[w+i] == z[i] {
		} else {
			return -1
		}
	}
	return w
}
