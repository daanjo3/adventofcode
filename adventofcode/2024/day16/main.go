package main

import (
	"fmt"
	"math"
	"slices"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day16",
		runMazeSimulation,
		c.PlaceholderFunc,
	)
}

type MazeBlock int

const (
	PATH MazeBlock = iota
	SOLID
	START
	END
)

func (mb MazeBlock) String() string {
	switch mb {
	case PATH:
		return "."
	case SOLID:
		return "#"
	case START:
		return "S"
	case END:
		return "E"
	default:
		panic("reached default case")
	}
}

func blockFromChar(char rune) MazeBlock {
	switch char {
	case '#':
		return SOLID
	case '.':
		return PATH
	case 'S':
		return START
	case 'E':
		return END
	default:
		panic(fmt.Sprintf("reached default case for char: %s", string(char)))
	}
}

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
	pos      c.Point
	facing   c.Ordinal
	pathCost int
}

func (r *Reindeer) clone() Reindeer {
	return Reindeer{
		maze:     r.maze,
		visited:  slices.Clone(r.visited),
		pos:      r.pos,
		facing:   r.facing,
		pathCost: r.pathCost,
	}
}

func (r *Reindeer) listMoves() []c.Ordinal {
	possible := []c.Ordinal{}
	for _, dir := range c.Ordinals {
		newPos := r.pos.Added(dir.ToPoint())
		if slices.ContainsFunc(r.visited, func(visitedPos c.Point) bool { return visitedPos == newPos }) {
			continue
		}
		block := r.maze.get(newPos)
		if block == SOLID {
			continue
		}
		possible = append(possible, dir)
	}
	return possible
}

func (r *Reindeer) moveCost(dir c.Ordinal) int {
	turningCost := 0
	if r.facing != dir {
		for r.facing != dir {
			r.facing = r.facing.TurnClockwise()
			turningCost += 1000
		}
		// Pretend we turned around counter clockwise
		if turningCost == 3000 {
			turningCost = 1000
		}
	}
	return turningCost + 1
}

func (r *Reindeer) move() Reindeer {
	moves := r.listMoves()
	if len(moves) == 0 {
		return Reindeer{pathCost: math.MaxInt}
	}

	fmt.Printf("Reindeer %v. Moves: %v\n", r.pos, moves)

	uberCheapestDeer := Reindeer{pathCost: math.MaxInt}
	for _, move := range moves {
		newDeer := r.clone()
		newDeer.pathCost += newDeer.moveCost(move)
		newDeer.pos = newDeer.pos.Added(move.ToPoint())
		newDeer.visited = append(newDeer.visited, newDeer.pos)
		newDeer.facing = move
		if newDeer.maze.get(newDeer.pos) == END {
			fmt.Printf("Reaching end with cost %v!\n", newDeer.pathCost)
			return newDeer
		}
		cheapestDeer := newDeer.move()
		if cheapestDeer.pathCost < uberCheapestDeer.pathCost {
			uberCheapestDeer = cheapestDeer
		}
	}
	return uberCheapestDeer
}

func (m *Maze) printWithPath(r Reindeer) {
	for y, row := range *m {
		for x, block := range row {
			if slices.ContainsFunc(r.visited, func(visitedPoint c.Point) bool { return visitedPoint == c.Point{X: x, Y: y} }) {
				fmt.Printf("O")
			} else {
				fmt.Print(block)
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func runMazeSimulation(inputfile string) {
	maze := Maze{}
	var reindeer Reindeer
	c.ReadLines(inputfile, func(line string) {
		mazeRow := []MazeBlock{}
		for _, char := range line {
			block := blockFromChar(char)
			mazeRow = append(mazeRow, block)
		}
		maze = append(maze, mazeRow)
	})

	for y, row := range maze {
		for x, block := range row {
			if block == START {
				reindeer = Reindeer{
					maze:     &maze,
					visited:  []c.Point{},
					pos:      c.Point{X: x, Y: y},
					facing:   c.EAST,
					pathCost: 0,
				}
			}
		}
	}

	cheapestDeer := reindeer.move()

	maze.printWithPath(cheapestDeer)

	fmt.Printf("The cheapest route to the end is %v\n", cheapestDeer.pathCost)
}
