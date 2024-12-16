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

type Reindeer struct {
	pos    c.Point
	facing c.Ordinal
}

func (r Reindeer) forward() {
	r.pos = r.pos.Added(r.facing.ToPoint())
}

func (r Reindeer) turnLeft() {
	if r.facing == c.NORTH {
		r.facing = c.WEST
	}
	if r.facing == c.WEST {
		r.facing = c.SOUTH
	}
	if r.facing == c.SOUTH {
		r.facing = c.EAST
	}
	if r.facing == c.EAST {
		r.facing = c.NORTH
	}
}

func (r Reindeer) turnRight() {
	if r.facing == c.NORTH {
		r.facing = c.EAST
	}
	if r.facing == c.EAST {
		r.facing = c.SOUTH
	}
	if r.facing == c.SOUTH {
		r.facing = c.WEST
	}
	if r.facing == c.WEST {
		r.facing = c.NORTH
	}
}

func move(matrix [][]rune, visited []string) int {
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
