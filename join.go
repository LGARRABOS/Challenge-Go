package piscine

func Join(strs []string, sep string) string {
	var z string
	for a := 0; a < len(strs); a++ {
		if a == len(strs)-1 {
			z = z + string(strs[a])
			break
		}
		z = z + string(strs[a])
		z = z + sep
	}
	return string(z)
}
