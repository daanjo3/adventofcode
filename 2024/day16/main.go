package main

import (
	"log"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day15",
		c.PlaceholderFunc,
		c.PlaceholderFunc,
	)
}

const START = 'S'
const END = 'E'

func move(matrix [][]rune, visited []string) {
	// forward (if possible)
	// turn right (if unvisited)
	// turn left (if unvisited)
}

func findMazeCost(inputfile string) {
	matrix := c.ReadRuneMatrix(inputfile)
	faced := [][]string{}

	startPos, err := c.MatrixFindRune(matrix, START)
	if err != nil {
		log.Fatalln(err)
	}
	startHeading := c.PointHeading{X: startPos.X, Y: startPos.Y, Facing: c.EAST}
	endPos, err := c.MatrixFindRune(matrix, END)
	if err != nil {
		log.Fatalln(err)
	}

}
