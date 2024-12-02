package common

import (
	"strconv"
	"strings"
)

func CountInt(arr []int, target int) int {
	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	return count
}

func ParseIntArray(line string) []int {
	var intArr = []int{}
	digits := strings.Fields(line)
	for _, digit := range digits {
		intArr = append(intArr, Must(strconv.Atoi(digit)))
	}
	return intArr
}

func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func Sum(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}
