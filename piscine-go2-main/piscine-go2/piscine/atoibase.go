package piscine

func indexInBase(r rune, base string) int {
	for i, br := range base {
		if br == r {
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0
	for _, r := range s {
		digit := indexInBase(r, base)
		if digit == -1 {
			// invalid character for this base, return 0 per spec
			return 0
		}
		result = result*baseLen + digit
	}
	return result
}
