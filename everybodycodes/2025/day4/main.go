package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day4", CalibrateMill1, CalibrateMill2)
}

func CalibrateMill1(filename string) {
	gears := readInput(filename)
	turns := 2025.0
	totalRelation := calculateTotalRelation(gears)
	fmt.Printf("The last gear will make %d rotations.\n", int(totalRelation*turns))
}

func CalibrateMill2(filename string) {
	gears := readInput(filename)
	turns := 10000000000000.0
	totalRelation := calculateTotalRelation(gears)
	fmt.Printf("The first gear will need to make %d rotations.\n", int(math.Ceil(turns/totalRelation)))
}

func calculateTotalRelation(gears []float64) float64 {
	totalRelation := 1.0
	for i := 0; i < len(gears)-1; i++ {
		relative := gears[i] / gears[i+1]
		totalRelation *= relative
	}
	return totalRelation
}

func readInput(filename string) []float64 {
	var result []float64
	lib.ReadLines(filename, func(line string, _ int) {
		result = append(result, float64(lib.Must(strconv.Atoi(line))))
	})
	return result
}
