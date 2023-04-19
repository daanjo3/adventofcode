package main

import (
	"fmt"
	"strings"

	"github.com/daanjo3/adventofcode2022/helper"
)

func unique(comp string) string {
	unique := ""
	for _, item := range comp {
		if !strings.ContainsRune(unique, item) {
			unique += string(item)
		}
	}
	return unique
}

func intersect(coll1 string, coll2 string) string {
	intersected := ""
	for _, item := range coll2 {
		if strings.ContainsRune(coll1, item) {
			intersected += string(item)
		}
	}
	return intersected
}

func toValue(item rune) int {
	if item >= 97 {
		return int(item) - 96
	}
	return int(item) - 38
}

// WTF Go ?!
func getFirstRune(str string) rune {
	for _, c := range str {
		return c
	}
	panic("This should not be reached")
}

func main() {
	rucksacks := helper.ReadLines("input.txt")
	total := 0
	for i := 0; i < len(rucksacks); i += 3 {
		unique1 := unique(rucksacks[i])
		unique2 := unique(rucksacks[i+1])
		unique3 := unique(rucksacks[i+2])
		badge := intersect(intersect(unique1, unique2), unique3)
		if len(badge) != 1 {
			panic("Badge: " + badge)
		}
		total += toValue(getFirstRune(badge))
	}
	fmt.Println(total)
}
