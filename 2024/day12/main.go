package main

import (
	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day12",
		calculateFenchCosts,
		c.PlaceholderFunc,
	)
}

type Plot struct {
	Point        c.Point
	PerimiterVal int
	Region       int
	Type         rune
}

func calculateFenchCosts(inputfile string) {
	plots := [][]Plot{}

	matrix := c.ReadRuneMatrix(inputfile)
	for y, row := range matrix {
		plotsY := []Plot{}
		for x, char := range row {
			point := c.Point{X: x, Y: y}
			pval := 0
			c.ScanNeighbors(matrix, point, func(nval rune, npos c.Point) {
				if nval != char {
					pval++
				}
			})
			plotsY = append(plotsY, Plot{
				Point: c.Point{X: x, Y: y},
			})
		}
		plots = append(plots, plotsY)
	}

}

// 0 => 1
// 1 => 2024
// 2024 => 20 24
// 20 => 2 0
// 24 => 2 4
// 2 => 4048 => 40 48 => 4 0 4 8
// 0 new cycle
// 2 => 4048 => 40 48 => 4 0 4 8
// 4 => 8096 => 80 => 96 => 8 0 9 6
