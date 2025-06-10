package piscine

func ShoppingListSort(slice []string) []string {
	// Get the length of the slice
	n := len(slice)

	// Implement bubble sort to sort by string length
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare lengths of adjacent strings
			if len(slice[j]) > len(slice[j+1]) {
				// Swap elements if they are in the wrong order
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
