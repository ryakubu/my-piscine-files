package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	// Determine the initial direction
	dir := 0
	for i := 1; i < len(a); i++ {
		cmp := f(a[i-1], a[i])
		if cmp > 0 {
			dir = 1 // potential non-increasing
			break
		} else if cmp < 0 {
			dir = -1 // potential non-decreasing
			break
		}
		// if cmp == 0, continue to next pair to find direction
	}

	for i := 1; i < len(a); i++ {
		cmp := f(a[i-1], a[i])
		if dir == -1 { // non-decreasing expected
			if cmp > 0 {
				return false
			}
		} else if dir == 1 { // non-increasing expected
			if cmp < 0 {
				return false
			}
		} else {
			// all elements are equal, so sorted
		}
	}
	return true
}
