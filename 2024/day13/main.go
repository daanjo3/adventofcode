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
		findLeastAmountOfTokens,
		c.PlaceholderFunc,
	)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Extended Euclidean algorithm to find the coefficients (Bezout's identity)
func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return 1, 0, a // gcd(a, 0) = a
	} else {
		x1, y1, g := extendedGCD(b, a%b)
		return y1, x1 - (a/b)*y1, g
	}
}

func solve(ax, bx, px, ay, by, py int) (int, int) {
	fmt.Printf("Solving: %v = %v+%v and %v = %v+%v\n", px, ax, bx, py, ay, by)

	det := ax*by - ay*bx
	fmt.Println("determinant:", det)
	if det == 0 {
		return 0, 0
	}

	// Solve for a
	a := (px*by - py*bx) / det

	// Solve for b
	b := (ax*py - ay*px) / det

	if a < 0 || b < 0 {
		return 0, 0
	}

	fmt.Printf("A: %v, B: %v\n\n", a, b)
	if a > 100 || b > 100 {
		return 0, 0
	}
	return a, b
}

func findLeastAmountOfTokens(inputfile string) {
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
			a, b := solve(ax, bx, px, ay, by, py)
			cost += a*3 + b
		}
	})
	// Because the last empty line is not returned by readlines
	a, b := solve(ax, bx, px, ay, by, py)
	cost += a*3 + b

	fmt.Printf("In total I would be out of %v coins\n", cost)

}
