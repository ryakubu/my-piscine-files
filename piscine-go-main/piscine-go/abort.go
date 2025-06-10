package piscine

func Abort(a, b, c, d, e int) int {
	// Place the numbers in an array
	numbers := []int{a, b, c, d, e}

	// Manual sorting (simple bubble sort)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			if numbers[j] > numbers[j+1] {
				// Swap numbers[j] and numbers[j+1]
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	// The median is now at index 2
	return numbers[2]
}
