package lib

import "fmt"

type Point struct {
	X, Y int
}

type PointHeading struct {
	X, Y   int
	DX, DY int
}

func MakePointHeading(pos Point, heading Point) PointHeading {
	return PointHeading{
		X:  pos.X,
		Y:  pos.Y,
		DX: heading.X,
		DY: heading.Y,
	}
}

func (p PointHeading) String() string {
	return fmt.Sprintf("Point(x: %v, y: %v, dx: %v, dy: %v)", p.X, p.Y, p.DX, p.DY)
}

func (p Point) Added(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Point) Subtract(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) String() string {
	return fmt.Sprintf("Point(x: %v, y: %v)", p.X, p.Y)
}

func isInBounds[T any](matrix [][]T, point Point) bool {
	if point.Y < 0 || point.X < 0 {
		return false
	}
	if point.Y > len(matrix)-1 || point.X > len(matrix[0])-1 {
		return false
	}
	return true
}

type Ordinal int

const (
	NORTH Ordinal = iota
	EAST
	SOUTH
	WEST
)

var Ordinals = []Ordinal{
	NORTH,
	EAST,
	SOUTH,
	WEST,
}

func (o Ordinal) ToPoint() Point {
	switch o {
	case NORTH:
		return Point{X: 0, Y: -1}
	case EAST:
		return Point{X: 1, Y: 0}
	case SOUTH:
		return Point{X: 0, Y: 1}
	case WEST:
		return Point{X: -1, Y: 0}
	default:
		panic("unknown ordinal direction")
	}
}

func (o Ordinal) String() string {
	switch o {
	case NORTH:
		return "North"
	case EAST:
		return "East"
	case SOUTH:
		return "South"
	case WEST:
		return "West"
	default:
		panic("unknown ordinal direction")
	}
}

func (o Ordinal) TurnClockwise() Ordinal {
	switch o {
	case NORTH:
		return EAST
	case EAST:
		return SOUTH
	case SOUTH:
		return WEST
	case WEST:
		return NORTH
	default:
		panic("unknown ordinal direction")
	}
}

func (o Ordinal) TurnCounterClockwise() Ordinal {
	switch o {
	case NORTH:
		return WEST
	case WEST:
		return SOUTH
	case SOUTH:
		return EAST
	case EAST:
		return NORTH
	default:
		panic("unknown ordinal direction")
	}
}

func ScanNeighbors[T any](matrix [][]T, point Point, callback func(val *T, point Point, ord Ordinal)) {
	for _, ordinal := range Ordinals {
		neighborPos := point.Added(ordinal.ToPoint())
		if isInBounds(matrix, neighborPos) {
			neighborVal := matrix[neighborPos.Y][neighborPos.X]
			callback(&neighborVal, neighborPos, ordinal)
		}
	}
}
