package piscine

func ToLower(s string) string {
	i := []rune(s)
	for a := 0; a < len(s); a++ {
		if i[a] >= 65 && i[a] <= 90 {
			i[a] = i[a] + 32
		}
	}
	return string(i)
}
