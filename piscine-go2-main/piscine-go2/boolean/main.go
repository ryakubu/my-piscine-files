package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Function to print a string (character by character)
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Function to check if a number is even
func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func main() {
	// Define messages and constants
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	// Get the length of arguments passed to the program
	lengthOfArg := len(os.Args) - 1

	// Check if the number of arguments is even or odd and print the corresponding message
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
