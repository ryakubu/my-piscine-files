package piscine

// Compact removes all zero-value elements from the slice and returns the count of non-zero elements.
func Compact(ptr *[]string) int {
	// Create a new slice to hold the non-zero elements
	var compacted []string

	// Iterate through the original slice
	for _, v := range *ptr {
		if v != "" { // Check if the string is non-empty
			compacted = append(compacted, v)
		}
	}

	// Assign the new compacted slice back to the original pointer
	*ptr = compacted

	// Return the number of non-zero elements
	return len(compacted)
}
