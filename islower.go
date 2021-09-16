package piscine

func IsLower(s string) bool {
	i := []rune(s)
	for a := 0; a < len(s); a++ {
		if i[a] >= 'a' && i[a] <= 'z' {
		} else {
			return false
		}
	}
	return true
}
