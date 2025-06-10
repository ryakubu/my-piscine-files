package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0 // Return 0 for negative powers
	}
	if power == 0 {
		return 1 // Return 1 for any number raised to the power of 0
	}

	return nb * RecursivePower(nb, power-1) // Recursive case: nb * nb^(power-1)
}
