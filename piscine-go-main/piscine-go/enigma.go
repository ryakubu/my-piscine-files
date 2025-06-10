package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	*******c = tempA // a → c
	****d = tempC    // c → d
	*b = tempD       // d → b
	***a = tempB     // b → a
}
