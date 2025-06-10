package piscine

func ActiveBits(n int) int {
	count := 0

	// Loop through all the bits of the number
	for n > 0 {
		// If the least significant bit is 1, increment the count
		if n&1 == 1 {
			count++
		}
		// Right-shift the number by 1 to check the next bit
		n >>= 1
	}

	return count
}
