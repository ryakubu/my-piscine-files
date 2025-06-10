package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n" // If string has fewer than 3 characters, return a newline
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i]) // Add every third character to the result string
	}

	return result + "\n" // Return the result followed by newline
}
