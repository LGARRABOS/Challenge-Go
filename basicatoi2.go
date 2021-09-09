package piscine

func BasicAtoi2(s string) int {
	o := 0
	c := 0
	nb_word := []rune(s)
	for _, word := range nb_word {
		for i := '0'; i < word; i++ {
			c++
		}
		o = o*10 + c
		c = 0
	}
	return o
}
