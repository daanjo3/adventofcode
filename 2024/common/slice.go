package common

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
	digits := strings.Fields(line)
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

func MatrixRuneClone(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}
