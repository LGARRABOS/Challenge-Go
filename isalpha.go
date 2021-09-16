package piscine

func IsAlpha(s string) bool {
	i := []rune(s)
	for a := 0; a < len(s); a++ {
		if i[a] >= 'a' && i[a] <= 'z' || i[a] >= 'A' && i[a] <= 'Z' || i[a] >= '1' && i[a] <= '9' {
		} else {
			return false
		}
	}
	return true
}
