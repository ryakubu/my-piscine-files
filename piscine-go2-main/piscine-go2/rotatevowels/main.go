package main

import (
	"os"
)

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Stdout.Write([]byte("\n"))
		return
	}

	// Collect all vowels in order
	vowels := []byte{}
	for _, arg := range args {
		for i := 0; i < len(arg); i++ {
			if isVowel(arg[i]) {
				vowels = append(vowels, arg[i])
			}
		}
	}

	// Reverse vowels
	reverseIndex := len(vowels) - 1

	// Replace vowels in original positions with reversed ones
	for i, arg := range args {
		bytes := []byte(arg)
		for j := 0; j < len(bytes); j++ {
			if isVowel(bytes[j]) {
				bytes[j] = vowels[reverseIndex]
				reverseIndex--
			}
		}
		os.Stdout.Write(bytes)
		if i < len(args)-1 {
			os.Stdout.Write([]byte(" "))
		}
	}
	os.Stdout.Write([]byte("\n"))
}
