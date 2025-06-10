package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0 // Return 0 for negative numbers
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
		// Check for overflow by comparing the result to the previous value
		if result < 0 {
			return 0 // Return 0 if overflow occurs
		}
	}
	return result
}
