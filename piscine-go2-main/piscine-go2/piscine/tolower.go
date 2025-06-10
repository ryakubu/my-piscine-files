package piscine

func ToLower(s string) string {
	result := []rune{}
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r - 'A' + 'a'
		}
		result = append(result, r)
	}
	return string(result)
}
