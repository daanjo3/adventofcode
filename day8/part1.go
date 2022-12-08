package main

import (
	"fmt"
	"math"

	"github.com/daanjo3/adventofcode2022/helper"
)

func Init2DArray(yMax int, xMax int) [][]int {
	grid := make([][]int, yMax)
	for y := 0; y < yMax; y++ {
		grid[y] = make([]int, xMax)
	}
	return grid
}

func CreateGrid(inputfile string) [][]int {
	lines := helper.ReadLines(inputfile)
	grid := Init2DArray(len(lines), len(lines[0]))
	for y, line := range lines {
		for x, height := range line {
			grid[y][x] = int(height - '0')
		}
	}
	return grid
}

func calculateVisible(grid [][]int) [][]int {
	diffMap := Init2DArray(len(grid), len(grid[0]))
	for y, row := range grid {
		for x, _ := range row {
			if x == 0 || x == len(row)-1 || y == 0 || y == len(grid)-1 {
				diffMap[y][x] = 1
				continue
			}
			if isVisible(grid, y, x) {
				diffMap[y][x] = 1
			} else {
				diffMap[y][x] = 0
			}
		}
	}
	return diffMap
}

func isVisible(grid [][]int, y int, x int) bool {
	topDiff := maxDirection(grid, y, x, -1, 0)
	bottomDiff := maxDirection(grid, y, x, 1, 0)
	leftDiff := maxDirection(grid, y, x, 0, -1)
	rightDiff := maxDirection(grid, y, x, 0, 1)
	visible := grid[y][x] > Min([]int{topDiff, bottomDiff, leftDiff, rightDiff})
	return visible
}

func Min(values []int) int {
	least := math.MaxInt
	for _, value := range values {
		if value < least {
			least = value
		}
	}
	return least
}

func maxDirection(grid [][]int, y int, x int, dy int, dx int) int {
	max := math.MinInt
	for y < len(grid)-1 && y > 0 && x < len(grid[0])-1 && x > 0 {
		y += dy
		x += dx
		if grid[y][x] > max {
			max = grid[y][x]
		}
	}
	return max
}

func countNonZero(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, cel := range row {
			if cel > 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	grid := CreateGrid("input.txt")
	fmt.Println(grid)
	diffMap := calculateVisible(grid)
	fmt.Println(diffMap)
	nonZeroCount := countNonZero(diffMap)
	fmt.Println(nonZeroCount)
}
