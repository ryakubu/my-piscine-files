package piscine

// FindNextPrime returns the first prime number greater than or equal to nb
func FindNextPrime(nb int) int {
	// Check if nb itself is prime, otherwise increment until we find the next prime
	for !IsPrime(nb) {
		nb++
	}
	return nb
}
