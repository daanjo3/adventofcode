package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/2024/common"
)

/**
 * For each line of the ordering
 * Check if either left or right is in ordering array
 * If neither add both to end of array
 * If one, put the missing token directly in front or behind the other
 * If both,
 */

func main() {
	common.AdventCommand("day5", sortPrintFilesUnsolved, sortPrintFilesSolved)
}

func updateSortMap(sortMap [][]int, line string) [][]int {
	valuesStr := strings.Split(line, "|")

	return append(sortMap, []int{
		common.Must(strconv.Atoi(valuesStr[0])),
		common.Must(strconv.Atoi(valuesStr[1])),
	})
}

func findMapEntry(sortMap [][]int, i, j int) (bool, []int) {
	for _, entry := range sortMap {
		if (entry[0] == i && entry[1] == j) || (entry[0] == j && entry[1] == i) {
			return true, entry
		}
	}
	return false, nil
}

func parseLine(line string) []int {
	valuesStr := strings.Split(line, ",")
	values := []int{}
	for _, valueStr := range valuesStr {
		values = append(values, common.Must(strconv.Atoi(valueStr)))
	}
	return values
}

func sortPrintFilesUnsolved(inputfile string) {
	sortPrintFiles(inputfile, false)
}

func sortPrintFilesSolved(inputfile string) {
	sortPrintFiles(inputfile, true)
}

func sortPrintFiles(inputfile string, solve bool) {
	sortMap := [][]int{}
	sumMiddle := 0

	common.ReadLines(inputfile, func(line string) {
		if strings.Contains(line, "|") {
			sortMap = updateSortMap(sortMap, line)
		}
		if strings.Contains(line, ",") {
			values := parseLine(line)
			valuesCp := parseLine(line)
			fmt.Println("unsorted", values)
			sort.Slice(values, func(i, j int) bool {
				found, entry := findMapEntry(sortMap, values[i], values[j])
				if !found {
					return false
				}
				if entry[0] == values[i] {
					return true
				}
				return false
			})
			if solve {
				sumMiddle += values[len(values)/2]
			} else if reflect.DeepEqual(values, valuesCp) {
				sumMiddle += values[len(values)/2]
			}
		}
	})

	fmt.Printf("Sum of middle page numbers is %v\n", sumMiddle)
}
