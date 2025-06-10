package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := 99
	for first >= 1 {
		second := first - 1
		for second >= 0 {
			printTwoDigits(first)
			z01.PrintRune(' ')
			printTwoDigits(second)
			if !(first == 1 && second == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			second--
		}
		first--
	}
}

// printTwoDigits prints an integer as two digits (leading zero if needed)
func printTwoDigits(n int) {
	z01.PrintRune(rune('0' + n/10))
	z01.PrintRune(rune('0' + n%10))
}
