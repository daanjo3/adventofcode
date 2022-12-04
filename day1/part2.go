package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/daanjo3/adventofcode2022/helper"
)

func calculateElfCapacities(calories []string) []int {
	elfCapacity := []int{0}
	elfSum := 0
	for _, line := range calories {
		if line == "" {
			elfCapacity = append(elfCapacity, elfSum)
			elfSum = 0
		} else {
			parsed, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			elfSum += parsed
		}
	}
	return elfCapacity
}

func main() {
	calories := helper.ReadLines("input.txt")
	elfCapacity := calculateElfCapacities(calories)

	// Reverse sort the list
	sort.Sort(sort.Reverse(sort.IntSlice(elfCapacity)))

	// Sum the top 3
	total := elfCapacity[0] + elfCapacity[1] + elfCapacity[2]
	fmt.Println(total)
}
