package piscine

func ToUpper(s string) string {
	i := []rune(s)
	for a := 0; a < len(s); a++ {
		if i[a] >= 97 && i[a] <= 122 {
			i[a] = i[a] - 32
		}
	}
	return string(i)
}
