package piscine

func BasicJoin(elems []string) string {
	var z string
	for a := 0; a < len(elems); a++ {
		z = z + elems[a]
	}
	return string(z)
}
