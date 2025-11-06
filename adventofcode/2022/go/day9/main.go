package main

import (
	"fmt"
	"os"
)

func getInput() string {
	args := os.Args[1:]
	if len(args) == 2 {
		return args[1]
	}
	return "input.txt"
}

func getPart() string {
	args := os.Args[1:]
	if len(args) == 1 {
		return args[0]
	}
	return "1"
}

func main() {
	part := getPart()
	input := getInput()

	if part == "1" {
		fmt.Println("Running part 1")
		RunPart1(input)
	}
	if part == "2" {
		fmt.Println("Part 2 not yet implemented")
	}

}
