package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day13",
		func(inputfile string) { findLeastAmountOfTokens(inputfile, 0) },
		func(inputfile string) { findLeastAmountOfTokens(inputfile, 10000000000000) },
	)
}

func solve(px, py, ax, ay, bx, by int, offset int) int {
	priceX := px + offset
	priceY := py + offset

	det := ax*by - ay*bx

	a := (priceX*by - priceY*bx) / det
	b := (ax*priceY - ay*priceX) / det

	if ax*a+bx*b == priceX && ay*a+by*b == priceY {
		return a*3 + b
	}
	return 0
}

func findLeastAmountOfTokens(inputfile string, offset int) {
	cost := 0
	buttonPattern := regexp.MustCompile(`Button [A,B]: X\+(?P<X>\d+), Y\+(?P<Y>\d+)`)
	prizePattern := regexp.MustCompile(`Prize: X=(?P<X>\d+), Y=(?P<Y>\d+)`)

	ax, ay := 0, 0
	bx, by := 0, 0
	px, py := 0, 0
	c.ReadLines(inputfile, func(line string) {
		if strings.Contains(line, "Button") {
			matches := buttonPattern.FindStringSubmatch(line)
			if strings.ContainsRune(line, 'A') {
				ax = c.Must(strconv.Atoi(matches[1]))
				ay = c.Must(strconv.Atoi(matches[2]))
			} else {
				bx = c.Must(strconv.Atoi(matches[1]))
				by = c.Must(strconv.Atoi(matches[2]))
			}
		}
		if strings.Contains(line, "Prize") {
			matches := prizePattern.FindStringSubmatch(line)
			px = c.Must(strconv.Atoi(matches[1]))
			py = c.Must(strconv.Atoi(matches[2]))
		}
		if len(line) == 0 {
			cost += solve(px, py, ax, ay, bx, by, offset)
		}
	})
	// Because the last empty line is not returned by readlines
	cost += solve(px, py, ax, ay, bx, by, offset)

	fmt.Printf("In total I would be out of %v coins\n", cost)

}
