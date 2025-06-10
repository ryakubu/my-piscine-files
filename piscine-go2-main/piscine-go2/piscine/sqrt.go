package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0 // Return 0 for negative numbers
	}

	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i // Return i if i^2 equals nb
		}
	}
	return 0 // Return 0 if no integer square root is found
}
