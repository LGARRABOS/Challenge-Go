package piscine

func BasicAtoi(s string) int {
	o := 0
	c := 0
	for _, word := range s {
		for i := '0'; i < word; i++ {
			c++
		}
		o = o*10 + c
		c = 0
	}
	return o
}
