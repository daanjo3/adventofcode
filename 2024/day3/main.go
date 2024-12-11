package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	x "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	x.AdventCommand("day3", parseUncorrupted, parseUncorruptedDoing)
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

func firstOrMax(values []int) int {
	if values == nil {
		return math.MaxInt
	}
	return values[0]
}

func parseUncorruptedDoing(inputfile string) {
	patternMul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	patternDont := regexp.MustCompile(`don't\(\)`)
	patternDo := regexp.MustCompile(`do\(\)`)

	total := 0
	enabled := true
	x.ReadLines(inputfile, func(line string) {
		remainder := line
		for {
			muls := patternMul.FindStringIndex(remainder)
			mul := firstOrMax(muls)
			dont := firstOrMax(patternDont.FindStringIndex(remainder))
			do := firstOrMax(patternDo.FindStringIndex(remainder))

			if do == mul && mul == dont {
				break
			}

			lowesti := 0
			if enabled {
				if mul < dont {
					op := remainder[muls[0]:muls[1]]
					total += handleOperation(op)
					lowesti = mul + 4
				} else {
					enabled = false
					lowesti = dont + 4
				}
			} else {
				if do == math.MaxInt {
					break
				}
				enabled = true
				lowesti = do + 4
			}

			remainder = remainder[lowesti:]
		}

	})

	fmt.Printf("The sum of valid operation outcomes is %v.\n", total)
}
