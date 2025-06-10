package piscine

func ReverseMenuIndex(menu []string) []string {
	// Create a new slice with the same length as the original
	reversed := make([]string, len(menu))

	// Loop over the original slice from the end to the start
	for i := 0; i < len(menu); i++ {
		// Fill the reversed slice starting from the last index
		reversed[i] = menu[len(menu)-1-i]
	}

	return reversed
}
