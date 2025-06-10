package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)
	// Iterate only up to the middle of the slice
	for i := 0; i < length/2; i++ {
		// Swap the element at index i with the element at index (length - 1 - i)
		podium[i], podium[length-1-i] = podium[length-1-i], podium[i]
	}
	return podium
}
