package main

import (
	"fmt"
	"slices"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day8", CountCenterCrosses, CountLineCrosses)
}

type Art struct {
	size  int
	lines []Line
}

func (a *Art) Add(line Line) int {
	knots := 0
	for _, l := range a.lines {
		if l.Crosses(a.size, line) {
			knots++
		}
	}
	a.lines = append(a.lines, line)
	return knots
}

type Line []int

func (l Line) Crosses(artSize int, other Line) bool {
	if other[0]-other[1] == 1 || other[1]-other[0] == 1 {
		return false
	}
	if slices.Contains(l, other[0]) || slices.Contains(l, other[1]) {
		return false
	}
	crosses := func(start, stop int) bool {
		count := 0
		for i := start; i <= stop; i++ {
			if i == artSize+1 {
				i = 1
			}
			if slices.Contains(l, i) {
				count++
			}
		}
		return count == 1
	}
	return crosses(other[0], other[1]) || crosses(other[1], other[0])
}

func CountLineCrosses(inputfile string) {
	numbers, max := parseStringArt(inputfile)
	art := Art{size: max, lines: []Line{}}
	sumKnots := 0
	for i := 0; i < len(numbers)-1; i++ {
		knots := art.Add(Line{numbers[i], numbers[i+1]})
		sumKnots += knots
	}
	fmt.Printf("Created %d knots in total\n", sumKnots)
}

func CountCenterCrosses(inputfile string) {
	numbers, max := parseStringArt(inputfile)
	count := 0
	for i := 0; i < len(numbers)-1; i++ {
		fmt.Println(numbers[i], numbers[i+1], max/2)
		if numbers[i+1] == numbers[i]+(max/2) || numbers[i+1] == numbers[i]-(max/2) {
			count++
		}
	}
	fmt.Printf("The line crosses the center %d times\n", count)
}

func parseStringArt(inputfile string) ([]int, int) {
	numbers := lib.ParseIntArray(lib.ReadLine(inputfile))
	max := 0
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return numbers, max
}
