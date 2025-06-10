package piscine

func Max(a []int) int {
	// If slice is empty, return 0
	if len(a) == 0 {
		return 0
	}

	// Initialize max with the first element
	max := a[0]

	// Iterate through the slice to find the maximum
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	return max
}
