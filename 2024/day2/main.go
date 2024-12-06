package main

import (
	"fmt"

	x "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	x.AdventCommand("day2", analyzeGraduality, analyzeGradualityLenient)
}

func isGradualTransition(report []int) bool {
	if report[0] == report[1] {
		return false
	}
	goingUp := report[0] < report[1]

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

		if diff < 1 || diff > 3 || diff == 0 {
			return false
		}
		levelPrev = level
	}
	return true
}

func isGradualTransitionV2(report []int) bool {
	if isGradualTransition(report) {
		return true
	}

	for i := range report {
		reportClone := make([]int, len(report))
		copy(reportClone, report)
		newReport := x.ArrRemoveInt(reportClone, i)
		if isGradualTransition(newReport) {
			return true
		}
	}
	return false
}

func analyzeGradualityLenient(inputfile string) {
	numSafeReports := 0

	x.ReadLines(inputfile, func(line string) {
		report := x.ParseIntArray(line)
		isSafe := isGradualTransitionV2(report)
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
		isSafe := isGradualTransition(report)
		if isSafe {
			numSafeReports++
		}
	})

	fmt.Printf("Of the reports, %v are reporting safe reactor behavior.", numSafeReports)
}
