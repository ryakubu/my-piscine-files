package piscine

import "github.com/01-edu/z01"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Use unsigned integer to handle absolute value safely
func printNbrBaseRec(n uint, base string, baseLen uint) {
	if n >= baseLen {
		printNbrBaseRec(n/baseLen, base, baseLen)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := uint(len(base))

	if nbr < 0 {
		z01.PrintRune('-')
		un := uint(-nbr)
		printNbrBaseRec(un, base, baseLen)
	} else {
		un := uint(nbr)
		printNbrBaseRec(un, base, baseLen)
	}
}
