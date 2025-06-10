package piscine

// RecursiveFactorial calculates the factorial of a given number
func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0 // Return 0 for negative numbers as factorial is not defined for negative integers
	}
	if nb == 0 {
		return 1 // Return 1 for 0! as it is defined as 1
	}

	// Check for potential overflow before multiplying
	if nb > 20 { // Factorials larger than 20! will overflow a 64-bit signed integer
		return 0
	}

	return nb * RecursiveFactorial(nb-1) // Recursive step
}
