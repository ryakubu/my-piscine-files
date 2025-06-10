package piscine

func TrimAtoi(s string) int {
	sign := 1
	foundSign := false
	foundDigit := false
	result := 0

	for _, r := range s {
		if !foundDigit {
			if r == '-' && !foundSign {
				sign = -1
				foundSign = true
				continue
			}
			if r >= '0' && r <= '9' {
				foundDigit = true
				result = int(r - '0')
				continue
			}
			if r == '+' && !foundSign {
				foundSign = true
				continue
			}
		} else {
			// already found digits, continue appending digits only
			if r >= '0' && r <= '9' {
				result = result*10 + int(r-'0')
			}
		}
	}

	if !foundDigit {
		return 0
	}
	return result * sign
}
