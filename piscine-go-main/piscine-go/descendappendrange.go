package piscine

func DescendAppendRange(max, min int) []int {
	// If max is less than or equal to min, return an empty slice
	if max <= min {
		return []int{}
	}

	// Initialize an empty slice
	var result []int

	// Iterate from max down to min + 1 and append each number
	for i := max; i > min; i-- {
		result = append(result, i)
	}

	// Return the result
	return result
}
