package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1 // Return -1 for negative indices
	}
	if index == 0 {
		return 0 // F(0) = 0
	}
	if index == 1 {
		return 1 // F(1) = 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2) // F(n) = F(n-1) + F(n-2)
}
