package main

import (
	"fmt"

	"github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	common.AdventCommand("day4", findXmasV1, findXmasV2)
}

func parseMatrix(inputfile string) [][]string {
	var matrix = [][]string{}
	common.ReadLines(inputfile, func(line string) {
		row := []string{}
		for _, r := range line {
			row = append(row, string(r))
		}
		matrix = append(matrix, row)
	})
	return matrix
}

func checkXmasV1(matrix [][]string, x int, y int, dir [2]int) bool {
	maxY := y + (3 * dir[1])
	if maxY > len(matrix)-1 || maxY < 0 {
		return false
	}
	maxX := x + (3 * dir[0])
	if maxX > len(matrix[y])-1 || maxX < 0 {
		return false
	}

	buffer := ""
	for range 4 {
		buffer += matrix[y][x]
		x += dir[0]
		y += dir[1]
	}
	return buffer == "XMAS"
}

func findXmasV1(inputfile string) {

	directions := [8][2]int{
		{1, 0},   // right
		{1, -1},  // right-down
		{0, -1},  // down
		{-1, -1}, // down-left
		{-1, 0},  // left
		{-1, 1},  // left-up
		{0, 1},   // up
		{1, 1},   // up-right
	}

	xmasFound := 0
	matrix := parseMatrix(inputfile)

	for y, row := range matrix {

		for x, c := range row {

			if c != "X" {
				continue
			}

			for _, direction := range directions {
				if checkXmasV1(matrix, x, y, direction) {
					xmasFound++
				}
			}

		}
	}

	fmt.Printf("Found the word XMAS %v times in the puzzle!\n", xmasFound)
}

func checkXmasV2(matrix [][]string, x int, y int) bool {
	if y+1 > len(matrix)-1 || y-1 < 0 {
		return false
	}
	if x+1 > len(matrix[y])-1 || x-1 < 0 {
		return false
	}

	corners := []string{
		matrix[y-1][x+1],
		matrix[y+1][x+1],
		matrix[y+1][x-1],
		matrix[y-1][x-1],
	}

	numM := common.CountString(corners, "M")
	numS := common.CountString(corners, "S")
	if numS != 2 || numM != 2 {
		return false
	}

	prevCorner := ""
	for _, corner := range corners {
		if corner == prevCorner {
			return true
		}
		prevCorner = corner
	}
	return false
}

func findXmasV2(inputfile string) {

	xmasFound := 0
	matrix := parseMatrix(inputfile)

	for y, row := range matrix {

		for x, c := range row {

			if c != "A" {
				continue
			}
			if checkXmasV2(matrix, x, y) {
				xmasFound++
			}

		}
	}

	fmt.Printf("Found the X-MAS %v times in the puzzle!\n", xmasFound)
}
