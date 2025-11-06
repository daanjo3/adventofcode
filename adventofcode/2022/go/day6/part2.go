package main

import (
	"fmt"
	"strings"

	"github.com/daanjo3/adventofcode2022/helper"
)

func anotherUnique(str string) string {
	unique := ""
	for _, val := range str {
		if !strings.ContainsRune(unique, val) {
			unique += string(val)
		}

	}
	return unique
}

func main() {
	input := helper.ReadLines("input.txt")
	buff := input[0]

	for i := 13; i < len(buff); i++ {
		// Slice string => 14 characters
		packet := buff[i-13 : i+1]
		if helper.Unique(packet) == packet {
			fmt.Println("packet", packet, len(packet))
			fmt.Println("char", i+1)
			break
		}
	}
}
