package main

import (
	"slices"

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

type MazeBlock int

const (
	PATH MazeBlock = iota
	SOLID
)

type Maze [][]MazeBlock

func (m *Maze) get(pos c.Point) MazeBlock {
	for y, row := range *m {
		for x, block := range row {
			if pos.Y == y && pos.X == x {
				return block
			}
		}
	}
	panic("invalid position")
}

type Reindeer struct {
	maze     *Maze
	visited  []c.Point
	pos      *c.Point
	facing   c.Ordinal
	pathCost int
}

func (r *Reindeer) listMoves() []c.Ordinal {
	possible := []c.Ordinal{}
	for _, dir := range c.Ordinals {
		newPos := r.pos.Added(dir.ToPoint())
		if slices.Contains(r.visited, newPos) {
			continue
		}
		if r.maze.get(newPos) == SOLID {
			continue
		}
		possible = append(possible, dir)
	}
	return possible
}
