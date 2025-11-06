package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day1", MyName, MomsName, DadsName)
}

func MyName(filename string) {
	instructions, names := ReadInput(filename)
	name := DetermineNameWithoutLoop(instructions, names)
	fmt.Printf("My.. Name.. Is.. '%s'!\n", name)
}

func MomsName(filename string) {
	instructions, names := ReadInput(filename)
	name := DetermineNameWithLoop(instructions, names)
	fmt.Printf("Moms.. Name.. Is.. '%s'!\n", name)
}

func DadsName(filename string) {
	instructions, names := ReadInput(filename)
	name := DetermineNameSwitching(instructions, names)
	fmt.Printf("Dads.. Name.. Is.. '%s'!\n", name)
}

type Instruction struct {
	Direction rune
	Distance  int
}

func NewInstructions(words []string) []Instruction {
	instructions := make([]Instruction, len(words))
	for i, word := range words {
		instructions[i] = NewInstruction(word)
	}
	return instructions
}

func NewInstruction(word string) Instruction {
	if len(word) < 2 {
		panic("invalid instruction length")
	}
	return Instruction{
		Direction: rune(word[0]),
		Distance:  lib.Must(strconv.Atoi(string(word[1:]))),
	}
}

func (i Instruction) String() string {
	return fmt.Sprintf("%c%d", i.Direction, i.Distance)
}

func ReadInput(inputfile string) ([]Instruction, []string) {
	var names []string
	var instructions []Instruction
	fmt.Println("Reading input from", inputfile)
	lib.ReadLines(inputfile, func(line string, index int) {
		if index == 0 {
			names = strings.Split(line, ",")
		}
		if index == 2 {
			instructions = NewInstructions(strings.Split(line, ","))
		}
	})
	return instructions, names
}

func DetermineNameWithoutLoop(instructions []Instruction, names []string) string {
	curIndex := 0
	for _, instr := range instructions {
		if instr.Direction == 'R' {
			curIndex = (curIndex + instr.Distance)
			if curIndex >= len(names) {
				curIndex = len(names) - 1
			}
		} else if instr.Direction == 'L' {
			curIndex = (curIndex - instr.Distance)
			if curIndex < 0 {
				curIndex = 0
			}
		} else {
			panic("invalid instruction: " + instr.String())
		}
	}
	return names[curIndex]
}

func DetermineNameWithLoop(instructions []Instruction, names []string) string {
	curIndex := 0
	for _, instr := range instructions {
		if instr.Direction == 'R' {
			curIndex = (curIndex + instr.Distance)
		} else if instr.Direction == 'L' {
			curIndex = (curIndex - instr.Distance)
		} else {
			panic("invalid instruction: " + instr.String())
		}
	}
	curIndex = curIndex % len(names)
	return names[curIndex]
}

func DetermineNameSwitching(instructions []Instruction, names []string) string {
	var otherIndex int
	namesCopy := make([]string, len(names))
	for _, instruction := range instructions {
		if instruction.Direction == 'R' {
			otherIndex = instruction.Distance % len(names)
		}
		if instruction.Direction == 'L' {
			otherIndex = (len(names) - (instruction.Distance % len(names))) % len(names)
		}

		copy(namesCopy, names)
		namesCopy[0] = names[otherIndex]
		namesCopy[otherIndex] = names[0]
		copy(names, namesCopy)
	}
	return names[0]
}
