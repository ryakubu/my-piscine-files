package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return 0 // in case string is empty
}
