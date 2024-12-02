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
