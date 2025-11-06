package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode2022/helper"
)

func getStackShort() []string {
	stack1 := "ZN"
	stack2 := "MCD"
	stack3 := "P"
	return []string{stack1, stack2, stack3}
}

func getStack() []string {
	return []string{
		"GFVHPS",
		"GJFBVDZM",
		"GMLJN",
		"NGZVDWP",
		"VRCB",
		"VRSMPWLZ",
		"THP",
		"QRSNCHZV",
		"FLGPVQJ",
	}
}

type operation struct {
	amount int
	origin int
	dest   int
}

func parseOp(line string) operation {
	split := strings.Split(line, " ")
	amount, err := strconv.Atoi(split[1])
	origin, err := strconv.Atoi(split[3])
	dest, err := strconv.Atoi(split[5])
	if err != nil {
		panic(err)
	}
	return operation{
		amount: amount,
		origin: origin,
		dest:   dest,
	}
}

// This can probably be cleaner with pointers and shit
func doOp(stacks []string, op operation) []string {
	ogStack := stacks[op.origin-1]
	destStack := stacks[op.dest-1]

	crates := strings.Clone(ogStack[len(ogStack)-op.amount:])
	ogStack = ogStack[:len(ogStack)-op.amount]
	destStack = destStack + crates
	stacks[op.origin-1] = ogStack
	stacks[op.dest-1] = destStack
	return stacks
}

func getTop(stacks []string) string {
	allTop := ""
	for _, stack := range stacks {
		allTop = allTop + string(stack[len(stack)-1])
	}
	return allTop
}

func main() {
	stacks := getStack()
	fmt.Println(stacks)
	for _, instruction := range helper.ReadLines("input.txt") {
		fmt.Println(instruction)
		op := parseOp(instruction)
		stacks = doOp(stacks, op)
	}
	fmt.Println(stacks)
	fmt.Println(getTop(stacks))
}
