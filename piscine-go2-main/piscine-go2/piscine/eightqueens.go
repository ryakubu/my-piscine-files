package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var solution [8]int

	var isSafe func(col, pos int) bool
	isSafe = func(col, pos int) bool {
		for i := 0; i < col; i++ {
			if solution[i] == pos || solution[i]-i == pos-col || solution[i]+i == pos+col {
				return false
			}
		}
		return true
	}

	var solve func(col int)
	solve = func(col int) {
		if col == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(solution[i] + '0'))
			}
			z01.PrintRune('\n')
			return
		}
		for pos := 1; pos <= 8; pos++ {
			if isSafe(col, pos) {
				solution[col] = pos
				solve(col + 1)
			}
		}
	}

	solve(0)
}
