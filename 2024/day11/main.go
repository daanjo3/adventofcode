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

func blink(stones []int) []int {
	newArr := []int{}
	for _, stone := range stones {
		newArr = append(newArr, transformStone(stone)...)
	}
	return newArr
}

func countStoneNew(stones []int, blinkTimes int, ch )

func countStonesN(stones []int, blinkTimes int, ch chan int) {
	for 

	for i := range blinkTimes {

		stones = blink(stones)


		for stone := stones {

		}

		fmt.Printf("Blinked for the %vth of %v time, now I got %v stones\n", i, blinkTimes, len(stones))

		if len(stones)%2 == 0 {
			newCh := make(chan int, 2)

			var left, right []int
			copy(left, stones[:len(stones)/2])
			copy(right, stones[len(stones)/2:])

			go countStonesN(left, blinkTimes-(i+1), newCh)
			go countStonesN(right, blinkTimes-(i+1), newCh)

			sum := 0
			for range 2 {
				sum += <-newCh
			}
			ch <- sum
			return
		}
	}
}

func countStones(inputfile string, blinkTimes int) {
	stones := c.ParseIntArray(c.ReadLine(inputfile))

	ch := make(chan int, len(stones))
	for _, stone := range stones {
		go countStonesN([]int{stone}, blinkTimes, ch)
	}

	sum := 0
	for val := range ch {
		sum += val
	}

	fmt.Printf("After blinking %v times I have %v stones!\n", blinkTimes, len(stones))
}

// 0 => 1
// 1 => 2024
// 2024 => 20 24
// 20 => 2 0
// 24 => 2 4
// 2 => 4048 => 40 48 => 4 0 4 8
// 0 new cycle
// 2 => 4048 => 40 48 => 4 0 4 8
// 4 => 8096 => 80 => 96 => 8 0 9 6
