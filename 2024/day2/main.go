package main

import (
	"fmt"

	x "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	x.AdventCommand("day2", analyzeGraduality, analyzeGradualityLenient)
}

func isGradualTransition(report []int, tokens int) bool {
	goingUp := true
	if report[0] == report[1] {
		return false
	}
	if report[0] > report[1] {
		goingUp = false
	}
	var levelPrev int
	for i, level := range report {
		if i == 0 {
			levelPrev = level
			continue
		}

		var diff int
		if goingUp {
			diff = level - levelPrev
		} else {
			diff = levelPrev - level
		}

		if diff < 1 || diff > 3 {
			if tokens > 0 {
				tokens--
				continue
			} else {
				return false
			}
		}
		levelPrev = level
	}
	return true
}

func analyzeGradualityLenient(inputfile string) {
	// Not yet working
	numSafeReports := 0

	x.ReadLines(inputfile, func(line string) {
		report := x.ParseIntArray(line)
		isSafe := isGradualTransition(report, 1)
		if isSafe {
			numSafeReports++
		}
	})

	fmt.Printf("Of the reports, %v are reporting safe reactor behavior using the problem dampener.", numSafeReports)
}

func analyzeGraduality(inputfile string) {
	numSafeReports := 0

	x.ReadLines(inputfile, func(line string) {
		report := x.ParseIntArray(line)
		isSafe := isGradualTransition(report, 0)
		if isSafe {
			numSafeReports++
		}
	})

	fmt.Printf("Of the reports, %v are reporting safe reactor behavior.", numSafeReports)
}
