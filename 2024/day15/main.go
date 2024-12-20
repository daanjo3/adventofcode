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
	objects []*Object
	robot   *Object
	width   int
	height  int
}

func (g *Grid) getObject(pos c.Point) *Object {
	for _, m := range g.objects {
		if *m.pos == pos {
			return m
		}
	}
	return nil
}

func (g *Grid) addRobot(pos *c.Point) {
	g.robot = &Object{grid: g, pos: pos, moveable: true, icon: '@'}
}

func (g *Grid) addBox(pos *c.Point) {
	g.objects = append(g.objects, &Object{grid: g, pos: pos, moveable: true, icon: 'O'})
}

func (g *Grid) addWall(pos *c.Point) {
	g.objects = append(g.objects, &Object{grid: g, pos: pos, moveable: false, icon: '#'})
}

func (g *Grid) print() {
	for y := range g.height {
		for x := range g.width {
			if g.robot.pos.X == x && g.robot.pos.Y == y {
				fmt.Print("@")
			} else {
				obj := g.getObject(c.Point{X: x, Y: y})
				if obj == nil {
					fmt.Print(".")
				} else if obj.moveable {
					fmt.Print("O")
				} else {
					fmt.Print("#")
				}
			}
		}
		fmt.Print("\n")
	}
}

type Object struct {
	grid     *Grid
	moveable bool
	pos      *c.Point
	icon     rune
}

func (m *Object) gps() int {
	return 100*(m.pos.Y+1) + m.pos.X + 1
}

func (m *Object) move(dir c.Ordinal) bool {
	newPos := m.pos.Added(dir.ToPoint())
	neighbor := m.grid.getObject(newPos)
	if neighbor != nil {
		if !neighbor.moveable {
			return false
		}
		if !neighbor.move(dir) {
			return false
		}
	}
	// fmt.Printf("Updating %s from %v to %v\n", string(m.icon), m.pos, newPos)
	m.pos = &newPos
	// m.grid.print()
	return true
}

func letRobotRun(inputfile string) {
	parsingMaze := true
	grid := Grid{
		width:   -1,
		height:  0,
		objects: []*Object{},
	}
	y := 0
	c.ReadLines(inputfile, func(line string) {
		if len(line) == 0 {
			parsingMaze = false
			return
		}
		if parsingMaze {
			if grid.width == -1 {
				grid.width = len(line)
			}
			grid.height++
			for x, char := range line {
				if char == '.' {
					continue
				}
				pos := c.Point{X: x, Y: y}
				if char == 'O' {
					grid.addBox(&pos)
				} else if char == '@' {
					grid.addRobot(&pos)
				} else if char == '#' {
					grid.addWall(&pos)
				} else {
					panic("Uncaught character")
				}
			}
			y++
		}
		if !parsingMaze {
			for _, char := range line {
				fmt.Println("Moving %s", string(char), c.NORTH, c.NORTH.ToPoint()) // North is actually south :)
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
				// grid.print()
			}
		}

	})

	// calculate GPS
	gpsSum := 0
	for _, box := range grid.objects {
		gpsSum += box.gps()
	}

	grid.print()

	fmt.Printf("The sum of all GPS scores is %v\n", gpsSum)
}
