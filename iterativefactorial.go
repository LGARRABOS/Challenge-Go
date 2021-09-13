package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 0; i < nb+1; i++ {
		result = result * i
	}
	return result
}
