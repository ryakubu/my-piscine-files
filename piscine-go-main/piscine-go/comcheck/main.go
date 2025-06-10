package main

import (
	"fmt"
	"os"
)

func main() {
	// Iterate over all the arguments passed to the program
	for _, arg := range os.Args[1:] {
		// Check if the argument matches any of the required strings
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			// If a match is found, print "Alert!!!" and exit
			fmt.Println("Alert!!!")
			return
		}
	}
}
