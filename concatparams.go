package piscine

func ConcatParams(args []string) string {
	var g string
	v := ('\n')
	for i := 0; i < len(args); i++ {
		g = g + args[i]
		if i != len(args)-1 {
			g = g + string(v)
		}
	}
	return g
}
