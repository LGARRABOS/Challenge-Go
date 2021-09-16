package piscine

func IsUpper(s string) bool {
	i := []rune(s)
	for a := 0; a < len(s); a++ {
		if i[a] >= 'A' && i[a] <= 'Z' {
		} else {
			return false
		}
	}
	return true
}
