package main

import "fmt"

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/daanjo3/adventofcode2022/helper"
// )

// type point struct {
// 	x int
// 	y int
// }

// func (p *point) hash() string {
// 	return strconv.Itoa(p.x) + "x" + strconv.Itoa(p.y)
// }

// func move(start *point, direction string) point {
// 	switch direction {
// 	case "U":
// 		return point{x: start.x, y: start.y + 1}
// 	case "D":
// 		return point{x: start.x, y: start.y - 1}
// 	case "R":
// 		return point{x: start.x + 1, y: start.y}
// 	case "L":
// 		return point{x: start.x - 1, y: start.y}
// 	}
// 	panic("Move reached end of method")
// }

// func getHorizontalStep(diff *point) string {
// 	if diff.x > 0 {
// 		return "R"
// 	}
// 	if diff.x < 0 {
// 		return "L"
// 	}
// 	return ""
// }

// func getVerticalStep(diff *point) string {
// 	if diff.y > 0 {
// 		return "U"
// 	}
// 	if diff.y < 0 {
// 		return "D"
// 	}
// 	return ""
// }

// func Abs(value int) int {
// 	if value > 0 {
// 		return value
// 	}
// 	return -value
// }

// func calculateDiff(head *point, tail *point) *point {
// 	return &point{x: head.x - tail.x, y: head.y - tail.y}
// }

// func calculateTailStep(head *point, tail *point) []string {
// 	diff := calculateDiff(head, tail)
// 	if Abs(diff.x) == 2 {
// 		if diff.y == 0 {
// 			return []string{getHorizontalStep(diff)}
// 		}
// 		return []string{getHorizontalStep(diff), getVerticalStep(diff)}
// 	}
// 	if Abs(diff.y) == 2 {
// 		if diff.x == 0 {
// 			return []string{getVerticalStep(diff)}
// 		}
// 		return []string{getVerticalStep(diff), getHorizontalStep(diff)}
// 	}
// 	return []string{}
// }

func Part1(filename string) {
	fmt.Println("Temporarily disabled")
}

// 	fmt.Println("Running part1 with " + filename)
// 	steps := helper.ReadLines(filename)

// 	head := &point{x: 0, y: 0}
// 	tail := &point{x: 0, y: 0}
// 	points := make(map[string]int)
// 	fmt.Println("head", head.hash())
// 	fmt.Println("tail", tail.hash())
// 	points[tail.hash()] = 1
// 	for _, step := range steps {
// 		stepArr := strings.Split(step, " ")
// 		direction := stepArr[0]
// 		amount, err := strconv.Atoi(stepArr[1])
// 		if err != nil {
// 			panic(err)
// 		}
// 		for i := 0; i < amount; i++ {
// 			fmt.Println("Stepping to " + direction)
// 			newHeadPoint := move(head, direction)
// 			head = &newHeadPoint
// 			tailSteps := calculateTailStep(head, tail)
// 			fmt.Println("Step causes tail steps", tailSteps)
// 			for _, tailStep := range tailSteps {
// 				newTailPoint := move(tail, tailStep)
// 				tail = &newTailPoint
// 			}
// 			points[tail.hash()] = 1
// 			// fmt.Println("head", head.hash())
// 			// fmt.Println("tail", tail.hash())
// 		}
// 	}

// 	count := 0
// 	for range points {
// 		count++
// 	}
// 	fmt.Println(points)
// 	fmt.Println("Unique fields", count)
// }
