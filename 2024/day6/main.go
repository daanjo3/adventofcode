package main

import (
	"fmt"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day6", analyzeGuardPath, loopGuardPath)
}

func findGuardStartingPos(matrix [][]rune) c.Point {
	pos, err := c.MatrixFindRune(matrix, '^')
	if err != nil {
		panic(err)
	}
	return pos
}

func isAboutToWalkOut(matrix [][]rune, guardPos c.Point, heading c.Point) bool {
	newPos := guardPos.Added(heading)
	if newPos.Y < 0 || newPos.Y >= len(matrix) {
		return true
	}
	if newPos.X < 0 || newPos.X >= len(matrix[newPos.Y]) {
		return true
	}
	return false
}

func nextHeading(heading c.Point) c.Point {
	if heading.X == 0 && heading.Y == -1 {
		return c.Point{X: 1, Y: 0}
	}
	if heading.X == 1 && heading.Y == 0 {
		return c.Point{X: 0, Y: 1}
	}
	if heading.X == 0 && heading.Y == 1 {
		return c.Point{X: -1, Y: 0}
	}
	if heading.X == -1 && heading.Y == 0 {
		return c.Point{X: 0, Y: -1}
	}
	panic("Invalid heading passed to function")
}

func analyzeGuardPath(inputfile string) {
	matrix := c.ReadRuneMatrix(inputfile)
	guardPos := findGuardStartingPos(matrix)
	heading := c.Point{X: 0, Y: -1}
	visited := make(map[c.Point]bool)
	visited[guardPos] = true
	fmt.Printf("Guard started at position %v\n", guardPos)

	for !isAboutToWalkOut(matrix, guardPos, heading) {
		newPos := guardPos.Added(heading)
		newPosVal := matrix[newPos.Y][newPos.X]
		if newPosVal == '#' {
			heading = nextHeading(heading)
		} else {
			guardPos = newPos
			visited[guardPos] = true
		}
		fmt.Printf("Guard moved to position %v, heading %v\n", guardPos, heading)
	}

	fmt.Printf("The guard visited %v cells in their patrol.\n", len(visited))
}

func loopGuardPath(inputfile string) {
	originalMatrix := c.ReadRuneMatrix(inputfile)
	walkedOut := false

	obstructed := make(map[c.Point]bool)

	loops := 0

	for !walkedOut {
		var matrix = c.MatrixRuneClone(originalMatrix)

		visitedUndirected := make(map[c.Point]bool)
		visited := make(map[c.PointHeading]bool)

		guardPos := findGuardStartingPos(matrix)
		heading := c.Point{X: 0, Y: -1}

		visited[c.MakePointHeading(guardPos, heading)] = true
		visitedUndirected[guardPos] = true
		fmt.Printf("Guard started at position %v\n", guardPos)

		brokeEarly := false

		for !isAboutToWalkOut(matrix, guardPos, heading) {
			newPos := guardPos.Added(heading)

			if _, ok := visitedUndirected[newPos]; ok {
				bearing := c.MakePointHeading(newPos, heading)
				fmt.Printf("Guard visited this position before: %v\n", bearing)

				// TODO look further ahead than just the next position
				if _, ok := visited[c.MakePointHeading(newPos, nextHeading(heading))]; ok {
					fmt.Printf("Guard did not walk this pos/heading before: %v\n", bearing)
					obstructionpos := newPos.Added(heading)
					if _, ok := obstructed[obstructionpos]; !ok {
						fmt.Printf("Could make new loop by obstructing point %v\n", obstructionpos)
						obstructed[obstructionpos] = true
						loops++
						brokeEarly = true
						break
					}
				}
			}

			newPosVal := matrix[newPos.Y][newPos.X]
			if newPosVal == '#' {
				heading = nextHeading(heading)
			} else {
				guardPos = newPos
				bearing := c.MakePointHeading(guardPos, heading)

				visitedUndirected[guardPos] = true
				visited[bearing] = true

				fmt.Printf("Guard moved to position %v\n", bearing)
			}
		}
		if !brokeEarly {
			walkedOut = true
		}
	}

	fmt.Printf("The guard could be looped on %v different ways.\n", loops)
}
