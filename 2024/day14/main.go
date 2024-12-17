package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func getSize(isSample bool) (int, int) {
	if isSample {
		return 11, 7
	}
	return 101, 103
}

func main() {
	c.AdventCommand("day14",
		findTheChillQuadrant,
		c.PlaceholderFunc,
	)
}

type Robot struct {
	Position c.Point
	Velocity c.Point
}

func maybeTeleport(pos c.Point, width, height int) c.Point {
	newX, newY := pos.X, pos.Y
	if pos.X < 0 {
		newX = width + pos.X
	}
	if pos.Y < 0 {
		newY = height + pos.Y
	}
	if pos.X > width-1 {
		newX = pos.X - width
	}
	if pos.Y > height-1 {
		newY = pos.Y - height
	}
	return c.Point{X: newX, Y: newY}
}

func simulate(robot *Robot, cycles, width, height int) {
	for range cycles {
		newPosX := robot.Position.X + robot.Velocity.X
		newPosY := robot.Position.Y + robot.Velocity.Y
		robot.Position = maybeTeleport(c.Point{X: newPosX, Y: newPosY}, width, height)
	}
}

func findTheChillQuadrant(inputfile string) {
	robotPattern := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	cycles := 100
	width, height := getSize(strings.Contains(inputfile, "sample"))

	robots := []Robot{}
	c.ReadLines(inputfile, func(line string) {
		matches := robotPattern.FindStringSubmatch(line)
		robot := Robot{
			Position: c.Point{X: c.Must(strconv.Atoi(matches[1])), Y: c.Must(strconv.Atoi(matches[2]))},
			Velocity: c.Point{X: c.Must(strconv.Atoi(matches[3])), Y: c.Must(strconv.Atoi(matches[4]))},
		}
		simulate(&robot, cycles, width, height)
		robots = append(robots, robot)
	})

	quadrants := [][]int{
		{0, width / 2, 0, height / 2},
		{0, width / 2, height/2 + 1, height},
		{width/2 + 1, width, 0, height / 2},
		{width/2 + 1, width, height/2 + 1, height},
	}

	safetyFactor := 1
	for _, quadrant := range quadrants {
		numRobots := countRobots(robots, quadrant)
		safetyFactor *= numRobots
	}

	fmt.Printf("The quandrant safety factor is %v\n", safetyFactor)
}

func countRobots(robots []Robot, quadrant []int) int {
	count := 0
	for _, robot := range robots {
		if robot.Position.X >= quadrant[0] &&
			robot.Position.X < quadrant[1] &&
			robot.Position.Y >= quadrant[2] &&
			robot.Position.Y < quadrant[3] {
			count++
		}
	}
	return count
}
