package piscine

func BasicAtoi2(s string) int {
	o := 0
	c := 0
	for _, word := range s {
		if word >= 48 && word <= 57 {
			for i := '0'; i < word; i++ {
				c++
			}
			o = o*10 + c
			c = 0
		} else {
			o = 0
			break
		}
	}
	return o
}
