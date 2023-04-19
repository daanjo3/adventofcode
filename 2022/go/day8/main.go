package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	part := args[0]
	input := args[1]
	if part == "1" {
		Part1(input)
	} else if part == "2" {
		Part2(input)
	} else {
		fmt.Println("Part " + part + " is not known")
	}
}
