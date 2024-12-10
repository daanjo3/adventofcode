package main

import (
	"fmt"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day10",
		sumTrails,
		c.PlaceholderFunc,
	)
}

func listNeighbors(matrix [][]int, curPos c.Point) []c.Point {
	points := []c.Point{}
	if curPos.X-1 >= 0 {
		points = append(points, c.Point{X: curPos.X - 1, Y: curPos.Y})
	}
	if curPos.X+1 < len(matrix[curPos.Y]) {
		points = append(points, c.Point{X: curPos.X + 1, Y: curPos.Y})
	}
	if curPos.Y-1 >= 0 {
		points = append(points, c.Point{X: curPos.X, Y: curPos.Y - 1})
	}
	if curPos.Y+1 < len(matrix) {
		points = append(points, c.Point{X: curPos.X, Y: curPos.Y + 1})
	}
	return points
}

func findTrails(matrix [][]int, curPos c.Point, height int) []c.Point {
	// fmt.Printf("Moved to %s, h=%v\n", curPos, height)
	if height == 9 {
		return []c.Point{curPos}
	}
	trails := []c.Point{}
	for _, nextPoint := range listNeighbors(matrix, curPos) {
		nextVal := matrix[nextPoint.Y][nextPoint.X]
		if nextVal == height+1 {
			trails = append(trails, findTrails(matrix, nextPoint, height+1)...)
		}
	}
	return trails
}

func sumTrails(inputfile string) {
	matrix := c.ReadIntMatrix(inputfile)
	sumTrailScores := 0
	sumTrailRatings := 0
	for y, row := range matrix {
		for x, val := range row {
			if val == 0 {
				trailhead := c.Point{X: x, Y: y}
				trails := findTrails(matrix, trailhead, 0)
				trailRating := len(trails)
				trailScore := len(c.ArrUniquePoint(trails))
				// fmt.Printf("Trailhead %s found %v trailtops\n", trailhead, trails)
				sumTrailScores += trailScore
				sumTrailRatings += trailRating
			}
		}
	}
	fmt.Printf("There are %v unique trail-head-top routes and the sum of rating is %v.\n", sumTrailScores, sumTrailRatings)
}
