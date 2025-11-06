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

func main() {
	rucksacks := helper.ReadLines("input.txt")
	fmt.Println(len(rucksacks))
	total := 0
	for _, rucksack := range rucksacks {
		c1 := rucksack[0:(len(rucksack) / 2)]
		c2 := rucksack[len(rucksack)/2:]
		c1Unique := unique(c1)
		c2Unique := unique(c2)

		for _, item := range intersect(c1Unique, c2Unique) {
			total += toValue(item)
		}
	}
	fmt.Println(total)
}
