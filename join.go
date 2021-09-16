package piscine

func Join(strs []string, sep string) string {
	var i string
	for a := 0; a < len(strs); a++ {
		if a == len(strs)-1 {
			i = i + string(strs[a])
			break
		}
		i = i + string(strs[a])
		i = i + sep
	}
	return string(i)
}
