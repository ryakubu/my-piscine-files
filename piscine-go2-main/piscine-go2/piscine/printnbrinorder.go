package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Handle 0 explicitly
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Count digits frequency (0-9)
	digitsCount := [10]int{}
	for n > 0 {
		digit := n % 10
		digitsCount[digit]++
		n /= 10
	}

	// Print digits in ascending order according to their count
	for digit := 0; digit <= 9; digit++ {
		for count := 0; count < digitsCount[digit]; count++ {
			z01.PrintRune(rune('0' + digit))
		}
	}
}
