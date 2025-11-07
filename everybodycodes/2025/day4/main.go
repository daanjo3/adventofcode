package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day4", CalibrateMill1, CalibrateMill2, CalibrateMill3)
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

func readInput(filename string) []float64 {
	var result []float64
	lib.ReadLines(filename, func(line string, _ int) {
		result = append(result, float64(lib.Must(strconv.Atoi(line))))
	})
	return result
}

func calculateTotalRelation(gears []float64) float64 {
	totalRelation := 1.0
	for i := 0; i < len(gears)-1; i++ {
		relative := gears[i] / gears[i+1]
		totalRelation *= relative
	}
	return totalRelation
}

func CalibrateMill3(filename string) {
	gears := readInput2(filename)
	turns := 100.0
	totalRelation := calculateTotalRelation2(gears)
	fmt.Printf("The last gear will make %d rotations.\n", int(totalRelation*turns))
}

func calculateTotalRelation2(gears []Gear) float64 {
	totalRelation := 1.0
	for i := 0; i < len(gears)-1; i++ {
		relative := gears[i].Back / gears[i+1].Front
		totalRelation *= relative
	}
	return totalRelation
}

type Gear struct {
	Front float64
	Back  float64
}

func (g Gear) String() string {
	if g.Front == 0 {
		return strconv.FormatFloat(g.Back, 'f', -1, 64)
	}
	if g.Back == 0 {
		return strconv.FormatFloat(g.Front, 'f', -1, 64)
	}
	return fmt.Sprintf("%d|%d", int(g.Front), int(g.Back))
}

func readInput2(filename string) []Gear {
	result := []Gear{}
	first := true
	lib.ReadLines(filename, func(line string, _ int) {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			front := float64(lib.Must(strconv.Atoi(parts[0])))
			back := float64(lib.Must(strconv.Atoi(parts[1])))
			result = append(result, Gear{Front: front, Back: back})
		} else {
			if first {
				result = append(result, Gear{Front: 0, Back: float64(lib.Must(strconv.Atoi(line)))})
				first = false
			} else {
				result = append(result, Gear{Front: float64(lib.Must(strconv.Atoi(line))), Back: 0})
			}
		}
	})
	return result
}
