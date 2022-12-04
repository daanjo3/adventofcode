package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode2022/helper"
)

func makeRangeArray(boundsStr string) []int {
	boundsArr := strings.Split(boundsStr, "-")
	lower, err1 := strconv.Atoi(boundsArr[0])
	upper, err2 := strconv.Atoi(boundsArr[1])
	if err1 != nil || err2 != nil {
		panic("Could not parse range list")
	}
	rangeArr := make([]int, upper-lower+1)
	for i := range rangeArr {
		rangeArr[i] = lower + i
	}
	return rangeArr
}

func main() {
	pairs := helper.ReadLines("input.txt")
	count := 0
	for _, pair := range pairs {
		pairArr := strings.Split(pair, ",")
		if len(pairArr) != 2 {
			panic("Pair does not have 2 units")
		}
		range1 := makeRangeArray(pairArr[0])
		range2 := makeRangeArray(pairArr[1])
		intersected := helper.Intersect(range1, range2)
		if len(intersected) == len(range1) || len(intersected) == len(range2) {
			count++
		}
	}
	fmt.Println(count)
}
