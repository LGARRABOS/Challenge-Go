package piscine

func Capitalize(s string) string {
	i := []rune(s)
	if i[0] >= 'a' && i[0] <= 'z' {
		i[0] = i[0] - 32
	}
	for a := 1; a < len(s); a++ {
		if i[a-1] >= 'a' && i[a]-1 <= 'z' || i[a-1] >= 'A' && i[a-1] <= 'Z' || i[a-1] >= '0' && i[a-1] <= '9' {
			if i[a] >= 'A' && i[a] <= 'Z' {
				i[a] = i[a] + 32
			}
		} else if i[a] >= 'a' && i[a] <= 'z' {
			i[a] = i[a] - 32
		}
	}
	j := string(i)
	return j
}
