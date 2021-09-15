package piscine

func AlphaCount(s string) int {
	i := []rune(s)
	m := 0
	for a := 0; a <= len(s); a++ {
		if i[a] >= 'a' && i[a] <= 'z' && i[a] >= 'A' && i[a] <= 'Z' {
			m++
		}
	}
	return m
}
