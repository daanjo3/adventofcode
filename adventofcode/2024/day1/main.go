package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	x "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	x.AdventCommand("day1", calculateDistance, calculateSimilarity)
}

func readInput(inputfile string) ([]int, []int) {
	left := []int{}
	right := []int{}

	x.ReadLines(inputfile, func(line string) {
		ids := strings.Fields(line)
		left = append(left, x.Must(strconv.Atoi(ids[0])))
		right = append(right, x.Must(strconv.Atoi(ids[1])))
	})

	return left, right
}

func calculateSimilarity(inputfile string) {
	similarity := 0
	left, right := readInput(inputfile)

	for _, target := range left {
		similarity += (target * x.ArrCountInt(right, target))
	}

	fmt.Printf("The similarity in location IDs lists is %v!\n", similarity)
}

func calculateDistance(inputfile string) {
	distance := 0
	left, right := readInput(inputfile)

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		distance += x.Abs(left[i] - right[i])
	}

	fmt.Printf("The difference in location IDs is %v!\n", distance)
}
