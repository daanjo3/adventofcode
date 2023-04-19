package main

import (
	"fmt"

	"github.com/daanjo3/adventofcode2022/helper"
)

func main() {
	input := helper.ReadLines("input.txt")
	buff := input[0]

	for i := 3; i < len(buff); i++ {
		// Slice string => 4 characters
		packet := buff[i-3 : i+1]
		if len(helper.Unique(packet)) == 4 {
			fmt.Println(i + 1)
			break
		}
	}
}
