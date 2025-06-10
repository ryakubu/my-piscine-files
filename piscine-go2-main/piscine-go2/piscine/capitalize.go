package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	inWord := false

	for i, ch := range runes {
		if isAlnum(ch) {
			if !inWord {
				// Start of a new word: capitalize if letter
				if ch >= 'a' && ch <= 'z' {
					runes[i] = ch - ('a' - 'A')
				}
				inWord = true
			} else {
				// Inside a word: lowercase if letter
				if ch >= 'A' && ch <= 'Z' {
					runes[i] = ch + ('a' - 'A')
				}
			}
		} else {
			// Non-alphanumeric character ends a word
			inWord = false
		}
	}
	return string(runes)
}

// Helper function to check if rune is alphanumeric
func isAlnum(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'a' && ch <= 'z') ||
		(ch >= '0' && ch <= '9')
}
