package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		start := i * 3
		end := start + 3
		playerCards := deck[start:end]
		fmt.Printf("Player %d: %d, %d, %d\n", i+1, playerCards[0], playerCards[1], playerCards[2])
	}
}
