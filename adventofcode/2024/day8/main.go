package main

import (
	"fmt"
	"slices"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day8",
		func(f string) { countAntiNodes(f, false) },
		func(f string) { countAntiNodes(f, true) },
	)
}

func frequencies() []rune {
	frequencies := []rune{}
	for i := '0'; i <= '9'; i++ {
		frequencies = append(frequencies, rune(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		frequencies = append(frequencies, rune(i))
	}
	for i := 'a'; i <= 'z'; i++ {
		frequencies = append(frequencies, rune(i))
	}
	return frequencies
}

func inBounds(matrix [][]rune, p c.Point) bool {
	return p.Y >= 0 && p.Y < len(matrix) && p.X >= 0 && p.X < len(matrix[0])
}

func countAntiNodes(inputfile string, countHarmonics bool) {
	matrix := c.ReadRuneMatrix(inputfile)

	antiPoints := []c.Point{}

	for _, freq := range frequencies() {

		points := c.MatrixFindRunes(matrix, freq)
		if len(points) == 0 {
			continue
		}
		fmt.Printf("%s: %s\n", string(freq), points)

		pointsCp := make([]c.Point, len(points))
		for i, point := range points {
			copy(pointsCp, points)
			otherPoints := append(pointsCp[:i], pointsCp[i+1:]...)

			for _, other := range otherPoints {
				antiPoint := point.Subtract(other.Subtract(point))

				if inBounds(matrix, antiPoint) && !slices.Contains(antiPoints, antiPoint) {
					fmt.Printf("%s => %s\n", string(freq), antiPoint)
					antiPoints = append(antiPoints, antiPoint)
				}
			}
		}

		// fmt.Printf("%s => %v\n", string(freq), antiPointsFreq)
	}

	fmt.Printf("Found %v antipoints!\n", len(antiPoints))
}
