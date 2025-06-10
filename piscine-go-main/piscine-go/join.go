package piscine

func Join(strs []string, sep string) string {
	// If the slice is empty, return an empty string
	if len(strs) == 0 {
		return ""
	}

	// Initialize the result with the first string in the slice
	result := strs[0]

	// Loop through the remaining strings and append them with the separator
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}

	return result
}
