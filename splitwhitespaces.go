package piscine

func SplitWhiteSpaces(s string) []string {
	var g []string
	var m string
	w := []rune(s)
	for i := 0; i < len(s); i++ {
		if w[i] == 32 || w[i] == '\n' || w[i] == 9 {
			if w[i+1] == 32 {
			} else {
				g = append(g, m)
				m = ""
			}
		} else {
			m = m + string(w[i])
		}
	}
	g = append(g, m)
	return g
}
