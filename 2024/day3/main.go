package main

import (
	"fmt"
	"regexp"
	"strconv"

	x "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	x.AdventCommand("day3", parseUncorrupted, parseUncorruptedDoing)
}

func findEnabledRange(dos [][]int, donts [][]int) (int, int) {
	return 0, 0
}

func parseUncorruptedDoing(inputfile string) {
	// patternMul := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	patternDont := regexp.MustCompile(`don't\(\)`)
	patternDo := regexp.MustCompile(`do\(\)`)

	total := 0
	x.ReadLines(inputfile, func(line string) {
		donts := patternDont.FindAllStringIndex(line, -1)
		dos := patternDo.FindAllStringIndex(line, -1)
		// operations := patternMul.FindAllStringIndex(line, -1)
		// fmt.Printf("dos: %v\ndon'ts: %v\nops: %v\n", dos, donts, operations)
		fmt.Printf("dos: %v\ndon'ts: %v\n", dos, donts)
	})

	fmt.Printf("The sum of valid operation outcomes is %v.\n", total)
}

func handleOperation(op string) int {
	pattern := regexp.MustCompile(`[0-9]{1,3}`)
	outcome := 1
	for _, strVal := range pattern.FindAllString(op, 2) {
		outcome *= x.Must(strconv.Atoi(strVal))
	}
	return outcome
}

func parseUncorrupted(inputfile string) {
	pattern := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	total := 0
	x.ReadLines(inputfile, func(line string) {
		operations := pattern.FindAllString(line, -1)
		for _, op := range operations {
			total += handleOperation(op)
		}
	})

	fmt.Printf("The sum of valid operation outcomes is %v.\n", total)
}
