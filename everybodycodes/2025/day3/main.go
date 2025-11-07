package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day3", StackCrates)
}

func ReadInput(inputfile string) []int {
	numbers := []int{}
	for _, numstr := range strings.Split(lib.ReadLine(inputfile), ",") {
		numbers = append(numbers, lib.Must(strconv.Atoi(numstr)))
	}
	return numbers
}

func StackCrates(inputfile string) {
	numbers := ReadInput(inputfile)
	slices.Sort(numbers)
	slices.Reverse(numbers)
	curNum := numbers[0]
	setSize := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < curNum {
			curNum = numbers[i]
			setSize += curNum
		}
	}
	fmt.Printf("Found %d of size %d\n", curNum, setSize)
}
