package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	// Parse first number
	a, ok := atoi(args[0])
	if !ok {
		return
	}

	// Parse operator
	op := args[1]
	if op != "+" && op != "-" && op != "*" && op != "/" && op != "%" {
		return
	}

	// Parse second number
	b, ok := atoi(args[2])
	if !ok {
		return
	}

	// Handle division/modulo by zero
	if b == 0 {
		if op == "/" {
			os.Stdout.WriteString("No division by 0\n")
			return
		} else if op == "%" {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
	}

	// Perform operation with overflow checks
	var result int64
	var overflow bool
	switch op {
	case "+":
		result, overflow = addOverflow(a, b)
	case "-":
		result, overflow = subOverflow(a, b)
	case "*":
		result, overflow = mulOverflow(a, b)
	case "/":
		result = a / b
	case "%":
		result = a % b
	}

	if overflow {
		return
	}
	os.Stdout.WriteString(itoa(result) + "\n")
}

// Custom string to int64 conversion
func atoi(s string) (int64, bool) {
	var n int64
	var sign int64 = 1
	i := 0

	if len(s) == 0 {
		return 0, false
	}

	if s[0] == '-' {
		sign = -1
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int64(s[i] - '0')
		// Check for overflow
		if n > (1<<63-1-digit)/10 {
			return 0, false
		}
		n = n*10 + digit
	}

	n *= sign
	// Special case for -2^63
	if sign == -1 && n > 0 {
		return 0, false
	}
	return n, true
}

// Custom int64 to string conversion
func itoa(n int64) string {
	if n == 0 {
		return "0"
	}

	var buf [20]byte // Enough for 64-bit number
	i := len(buf) - 1
	negative := n < 0

	if negative {
		n = -n
	}

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i--
	}

	if negative {
		buf[i] = '-'
		i--
	}

	return string(buf[i+1:])
}

// Helper functions to check for overflow
func addOverflow(a, b int64) (int64, bool) {
	if b > 0 && a > (1<<63-1)-b {
		return 0, true
	}
	if b < 0 && a < (-1<<63)-b {
		return 0, true
	}
	return a + b, false
}

func subOverflow(a, b int64) (int64, bool) {
	if b < 0 && a > (1<<63-1)+b {
		return 0, true
	}
	if b > 0 && a < (-1<<63)+b {
		return 0, true
	}
	return a - b, false
}

func mulOverflow(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}
	result := a * b
	if a == result/b {
		return result, false
	}
	return 0, true
}
