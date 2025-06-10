package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, char := range str {
		result = append(result, int(char)) // Convert each character to its ASCII value and append it to the slice
	}
	return result
}
