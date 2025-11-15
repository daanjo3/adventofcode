package main

import (
	"fmt"
	"slices"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day10", TestDragonMoves)
}

type piece = rune

const (
	sheep  piece = 'S'
	dragon piece = 'D'
	empty  piece = '.'
)

var dragonMoves = []lib.Point{
	{X: 1, Y: 2},
	{X: 1, Y: -2},
	{X: -1, Y: 2},
	{X: -1, Y: -2},
	{X: 2, Y: 1},
	{X: 2, Y: -1},
	{X: -2, Y: 1},
	{X: -2, Y: -1},
}

type Board struct {
	grid [][]piece
}

func (b Board) isValidPos(pos lib.Point) bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}
	sizeX, sizeY := b.size()
	if pos.X > sizeX-1 || pos.Y > sizeY-1 {
		return false
	}
	return true
}

func (b Board) size() (int, int) {
	return len(b.grid[0]), len(b.grid)
}

func (b Board) getDragon() Dragon {
	for y, row := range b.grid {
		for x, char := range row {
			if char == dragon {
				return Dragon{board: &b, pos: lib.Point{X: x, Y: y}}
			}
		}
	}
	panic("Couldn't find dragon")
}

func (b Board) hasSheep(pos lib.Point) bool {
	return b.grid[pos.Y][pos.X] == sheep
}

type Dragon struct {
	board *Board
	pos   lib.Point
}

func (d Dragon) nextPositions() []lib.Point {
	positions := []lib.Point{}
	for _, move := range dragonMoves {
		newPos := d.pos.Added(move)
		if !d.board.isValidPos(newPos) {
			continue
		}
		positions = append(positions, newPos)
	}
	return positions
}

func (d Dragon) findAllPositions(iterMax int, visited *[]lib.Point) {
	for _, next := range d.nextPositions() {
		if slices.ContainsFunc(*visited, func(point lib.Point) bool {
			return point.Equals(next)
		}) {
			continue
		}
		*visited = append(*visited, next)
		if iterMax > 0 {
			Dragon{board: d.board, pos: next}.findAllPositions(iterMax-1, visited)
		}
	}
}

func NewBoard(grid [][]rune) Board {
	return Board{grid: grid}
}

func TestDragonMoves(inputfile string) {
	board := NewBoard(lib.ReadRuneMatrix(inputfile))
	visited := []lib.Point{}
	dragon := board.getDragon()
	numMoves := 4
	dragon.findAllPositions(numMoves, &visited)
	count := 0
	for _, pos := range visited {
		if board.hasSheep(pos) {
			count++
		}
	}
	fmt.Printf("Found %d sheeps after %d dragon moves", count, numMoves)
}
