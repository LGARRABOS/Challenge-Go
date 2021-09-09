package piscine

func BasicAtoi(s string) int {
	o := 0
	c := 0
	nb_rune := []rune(s)
	for _, word := range nb_rune {
		for i := '0'; i < word; i++ {
			c++
		}
		o = 0*10 + c
		c = 0
	}
	return o
}
