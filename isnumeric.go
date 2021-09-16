package piscine

func IsNumeric(s string) bool {
	i := []rune(s)
	for a := 0; a < len(s); a++ {
		if i[a] >= '0' && i[a] <= '9' {
		} else {
			return false
		}
	}
	return true
}
