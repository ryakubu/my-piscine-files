package main

import "github.com/01-edu/z01"

func main() {
	letter := 'z'
	for letter >= 'a' {
		z01.PrintRune(letter)
		letter--
	}
	z01.PrintRune('\n')
}
