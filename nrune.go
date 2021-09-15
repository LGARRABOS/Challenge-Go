package piscine

func NRune(s string, n int) rune {
	i := []rune(s)
	if n > 0 && n <= len(s) {
		return (i[n-1])
	} else {
		return 0
	}
}
