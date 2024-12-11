package main

import (
	"fmt"
	"strconv"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day11",
		func(inputfile string) { countStones(inputfile, 25) },
		func(inputfile string) { countStones(inputfile, 75) },
	)
}

func transformStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	asStr := strconv.Itoa(stone)
	if len(asStr)%2 == 0 {
		left := asStr[:len(asStr)/2]
		right := asStr[len(asStr)/2:]
		return []int{
			c.Must(strconv.Atoi(left)),
			c.Must(strconv.Atoi(right)),
		}
	}
	return []int{stone * 2024}
}

func countStones(inputfile string, blinkTimes int) {
	stones := c.ParseIntArray(c.ReadLine(inputfile))

	countMap := map[int]int{}

	for _, stone := range stones {
		val, ok := countMap[stone]
		if !ok {
			countMap[stone] = 1
		} else {
			countMap[stone] = val + 1
		}
	}

	for range blinkTimes {

		newCountMap := map[int]int{}

		for num, count := range countMap {
			if count == 0 {
				continue
			}

			transformed := transformStone(num)

			for _, newNum := range transformed {

				newNumCount, ok := newCountMap[newNum]
				if !ok {
					newCountMap[newNum] = count
				} else {
					newCountMap[newNum] = newNumCount + count
				}

			}

		}

		countMap = newCountMap

	}

	fmt.Println(countMap)

	sum := 0
	for _, count := range countMap {
		sum += count
	}

	fmt.Printf("After blinking %v times I have %v stones!\n", blinkTimes, sum)
}
