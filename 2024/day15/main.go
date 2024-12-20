package main

import (
	"fmt"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day15",
		letRobotRun,
		c.PlaceholderFunc,
	)
}

type Grid struct {
	boxes  []Moveable
	robot  Moveable
	width  int
	height int
}

func (g *Grid) getMoveable(pos c.Point) *Moveable {
	for _, m := range g.boxes {
		if m.pos == pos {
			return &m
		}
	}
	return nil
}

func (g *Grid) outOfBounds(pos c.Point) bool {
	return pos.X < 0 || pos.Y < 0 || pos.X > g.width-1 || pos.Y > g.height-1
}

func (g *Grid) addBox(pos c.Point) {
	g.boxes = append(g.boxes, Moveable{grid: g, pos: pos})
}

type Moveable struct {
	grid *Grid
	pos  c.Point
}

func (m *Moveable) gps() int {
	return 100*(m.pos.Y+1) + m.pos.X + 1
}

func (m *Moveable) move(dir c.Ordinal) bool {
	newPos := m.pos.Added(dir.ToPoint())
	neighbor := m.grid.getMoveable(newPos)
	if neighbor != nil {
		if !neighbor.move(dir) {
			return false
		}
	}
	if m.grid.outOfBounds(newPos) {
		return false
	}
	m.pos = newPos
	return true
}

func letRobotRun(inputfile string) {
	parsingMaze := true
	grid := Grid{
		width:  -1,
		height: -2, // account for ##### borders
		boxes:  []Moveable{},
	}
	y := 0
	c.ReadLines(inputfile, func(line string) {
		if parsingMaze {
			if grid.width == -1 {
				grid.width = len(line) - 1
			}
			grid.height++
			for x, char := range line {
				if char == 'O' {
					grid.addBox(c.Point{X: x, Y: y})
				}
				if char == '@' {
					grid.robot = Moveable{grid: &grid, pos: c.Point{X: x, Y: y}}
				}
			}
			y++
		}
		if len(line) == 0 {
			parsingMaze = false
			return
		}
		if !parsingMaze {
			for _, char := range line {
				if char == '^' {
					grid.robot.move(c.NORTH)
				} else if char == '>' {
					grid.robot.move(c.EAST)
				} else if char == 'v' {
					grid.robot.move(c.SOUTH)
				} else if char == '<' {
					grid.robot.move(c.WEST)
				} else {
					panic(fmt.Sprintf("Unknown character %v", char))
				}
			}
		}

	})

	for y := range grid.height + 2 {
		for x := range grid.width + 2 {
			if y == 0 || y == grid.height+1 {
				fmt.Print("#")
			}
			if x == 0 || x == grid.width+1 {
				fmt.Print("#")
			}
			if grid.robot.pos.X == x+1 && grid.robot.pos.Y == y+1 {
				fmt.Print("@")
			}
			if grid.getMoveable(c.Point{X: x, Y: y}) == nil {
				fmt.Print(".")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Print("\n")
	}

	// calculate GPS
	gpsSum := 0
	for _, box := range grid.boxes {
		gpsSum += box.gps()
	}

	fmt.Printf("The sum of all GPS scores is %v\n", gpsSum)
}
