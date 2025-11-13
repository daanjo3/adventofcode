package lib

import (
	"errors"
	"strconv"
	"strings"
)

func ArrCountInt(arr []int, target int) int {
	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	return count
}

func ArrCountString(arr []string, target string) int {
	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	return count
}

func ArrRemoveInt(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func ArrSum(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func ParseIntArray(line string) []int {
	var intArr = []int{}
	var digits []string
	if strings.Contains(line, ",") {
		digits = strings.Split(line, ",")
	} else {
		digits = strings.Fields(line)
	}
	for _, digit := range digits {
		intArr = append(intArr, Must(strconv.Atoi(digit)))
	}
	return intArr
}

func MatrixFindRune(matrix [][]rune, target rune) (Point, error) {
	for y, row := range matrix {
		for x, tile := range row {
			if tile == target {
				return Point{X: x, Y: y}, nil
			}
		}
	}
	return Point{}, errors.New("target not found in matrix")
}

func MatrixFindRunes(matrix [][]rune, target rune) []Point {
	found := []Point{}
	for y, row := range matrix {
		for x, tile := range row {
			if tile == target {
				found = append(found, Point{X: x, Y: y})
			}
		}
	}
	return found
}

func MatrixRuneClone(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func ReadRuneMatrix(inputfile string) [][]rune {
	matrix := [][]rune{}
	ReadLines(inputfile, func(s string, _ int) {
		matrix = append(matrix, []rune(s))
	})
	return matrix
}

func ReadIntMatrix(inputfile string) [][]int {
	matrix := [][]int{}
	ReadLines(inputfile, func(s string, _ int) {
		row := []int{}
		for _, v := range s {
			row = append(row, Must(strconv.Atoi(string(v))))
		}
		matrix = append(matrix, row)
	})
	return matrix
}

func ArrUnique[T comparable](slice []T) []T {
	unique := []T{}
	for _, v := range slice {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}
	return unique
}
