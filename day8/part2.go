package main

import (
	"fmt"

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

func calculateViewscore(grid [][]int) int {
	diffMap := Init2DArray(len(grid), len(grid[0]))
	best := 0
	for y, row := range grid {
		for x, _ := range row {
			if x == 0 || x == len(row)-1 || y == 0 || y == len(grid)-1 {
				diffMap[y][x] = 2
				continue
			}
			score := viewScore(grid, y, x)
			if score > best {
				best = score
			}
		}
	}
	return best
}

func viewScore(grid [][]int, y int, x int) int {
	topDiff := viewDirection(grid, y, x, -1, 0)
	bottomDiff := viewDirection(grid, y, x, 1, 0)
	leftDiff := viewDirection(grid, y, x, 0, -1)
	rightDiff := viewDirection(grid, y, x, 0, 1)
	return topDiff * bottomDiff * leftDiff * rightDiff
}

func viewDirection(grid [][]int, y int, x int, dy int, dx int) int {
	count := 0
	myHeight := grid[y][x]
	for y < len(grid)-1 && y > 0 && x < len(grid[0])-1 && x > 0 {
		y += dy
		x += dx
		count++
		if grid[y][x] >= myHeight {
			break
		}
	}
	return count
}

func main() {
	grid := CreateGrid("input.txt")
	fmt.Println(grid)
	score := calculateViewscore(grid)
	fmt.Println(score)
}
