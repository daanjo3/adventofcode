package main

import (
	"fmt"

	x "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	x.AdventCommand("day2", analyzeGraduality, analyzeGradualityLenient)
}

func isGradualTransition(report []int, tokens int) (bool, int) {
	if (report[0] == report[1]) || (report[0]-report[1] > 3) || (report[0]-report[1] < -3) {
		tokens = 0
		report = report[1:]
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

		if diff < 1 || diff > 3 {
			if tokens > 0 {
				tokens--
				if i == 1 {
					levelPrev = level
				}
				if i == (len(report) - 1) {
					return true, 0
				}
				continue
			} else {
				return false, diff
			}
		}
		levelPrev = level
	}
	return true, 0
}

func calculateDiffs(report []int) []int {
	diffs := []int{}
	prev := report[0]
	for _, level := range report[1:] {
		diffs = append(diffs, prev-level)
		prev = level
	}
	return diffs
}

func isAsc(diffs []int) bool {
	numPos := 0
	for _, val := range diffs {
		if val > 0 {
			numPos++
		}
	}
	// fmt.Println("numpos", numPos, len(diffs))
	return numPos > len(diffs)-numPos
}

func isGradualTransitionV2(report []int, goDeeper bool) bool {
	diffs := calculateDiffs(report)
	numErr := 0
	isAscending := isAsc(diffs)
	// fmt.Println("is ascending", isAscending)
	for _, diff := range diffs {
		if isAscending {
			if diff < 0 {
				numErr++
				continue
			}
		} else {
			if diff > 0 {
				numErr++
				continue
			}
		}
		if diff < -3 || diff > 3 || diff == 0 {
			numErr++
		}
	}
	if numErr == 0 {
		return true
	}
	// fmt.Println("diffs", diffs)
	// fmt.Println("num err", numErr)
	if numErr > 1 {
		return false
	}
	if !goDeeper {
		return false
	}
	for i := range report {
		reportClone := make([]int, len(report))
		copy(reportClone, report)
		newReport := x.Remove(reportClone, i)
		if isGradualTransitionV2(newReport, false) {
			return true
		}
	}
	return false
}

func analyzeGradualityLenient(inputfile string) {
	numSafeReports := 0

	x.ReadLines(inputfile, func(line string) {
		report := x.ParseIntArray(line)
		isSafeV1, _ := isGradualTransition(report, 1)
		isSafeV2 := isGradualTransitionV2(report, true)
		if isSafeV1 != isSafeV2 {
			fmt.Printf("report: %s, v1: %t, v2: %t\n", line, isSafeV1, isSafeV2)
		}
		// fmt.Printf("Report %s is safe: %t\n", line, isSafeV2)
		if isSafeV2 {
			numSafeReports++
		}
	})

	fmt.Printf("Of the reports, %v are reporting safe reactor behavior using the problem dampener.", numSafeReports)
}

func analyzeGraduality(inputfile string) {
	numSafeReports := 0

	x.ReadLines(inputfile, func(line string) {
		report := x.ParseIntArray(line)
		isSafe, _ := isGradualTransition(report, 0)
		if isSafe {
			numSafeReports++
		}
	})

	fmt.Printf("Of the reports, %v are reporting safe reactor behavior.", numSafeReports)
}
