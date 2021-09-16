package piscine

func IsPrintable(s string) bool {
	i := []rune(s)
	for a := 0; a < len(s); a++ {
		if i[a] >= 32 && i[a] <= 126 {
		} else {
			return false
		}
	}
	return true
}
