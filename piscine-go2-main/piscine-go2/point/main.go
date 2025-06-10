package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(p *point) {
	p.x = 6 * 7
	p.y = 3 * 7
}

func main() {
	var p point
	setPoint(&p)

	output := []rune{
		120, 32, 61, 32, 52, 50, 44, 32, 121, 32, 61, 32, 50, 49, 10,
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
