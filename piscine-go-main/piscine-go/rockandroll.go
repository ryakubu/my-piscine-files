package piscine

func RockAndRoll(n int) string {
	// Check for negative numbers first
	if n < 0 {
		return "error: number is negative\n"
	}

	// Check if the number is divisible by both 2 and 3
	if n%2 == 0 && n%3 == 0 {
		return "rock and roll\n"
	}

	// Check if the number is divisible by 2
	if n%2 == 0 {
		return "rock\n"
	}

	// Check if the number is divisible by 3
	if n%3 == 0 {
		return "roll\n"
	}

	// If none of the above conditions are met, it's non-divisible
	return "error: non divisible\n"
}
