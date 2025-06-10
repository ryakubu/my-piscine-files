package piscine

func Rot14(s string) string {
	result := []rune{}

	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			// shift within 'a' to 'z'
			shifted := r + 14
			if shifted > 'z' {
				shifted = shifted - 26
			}
			result = append(result, shifted)
		case 'A' <= r && r <= 'Z':
			// shift within 'A' to 'Z'
			shifted := r + 14
			if shifted > 'Z' {
				shifted = shifted - 26
			}
			result = append(result, shifted)
		default:
			// leave non-letters unchanged
			result = append(result, r)
		}
	}

	return string(result)
}
