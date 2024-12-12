package common

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

func (p Point) String() string {
	return fmt.Sprintf("Point(x: %v, y: %v)", p.X, p.Y)
}

func ScanNeighbors[T any](matrix [][]T, point Point, callback func(val T, point Point)) {
	directions := []Point{
		{X: 1, Y: 0},  // right
		{X: 0, Y: -1}, // down
		{X: -1, Y: 0}, // left
		{X: 0, Y: 1},  // up
	}

	for _, direction := range directions {
		neighborPos := point.Added(direction)
		neighborVal := matrix[neighborPos.Y][neighborPos.X]
		callback(neighborVal, neighborPos)
	}
}
