package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day17",
		run,
		c.PlaceholderFunc,
	)
}

func parseRegister(line string) int {
	fields := strings.Fields(line)
	return c.Must(strconv.Atoi(fields[len(fields)-1]))
}

func parseInstructions(line string) []int {
	fields := strings.Fields(line)
	nums := strings.Split(fields[len(fields)-1], ",")
	arr := []int{}
	for _, strval := range nums {
		arr = append(arr, c.Must(strconv.Atoi(strval)))
	}
	return arr
}

type Registers struct {
	a int
	b int
	c int
}

func bitXor(valA, valB int) int {
	return asBinary(valA, valB, func(binA, binB string) string {
		outcome := ""
		for i := range binA {
			// XOR (to be tested)
			if (binA[i] == '1' && binB[i] != '1') || (binB[i] == '1' && binA[i] != '1') {
				outcome += "1"
			} else {
				outcome += "0"
			}
		}
		return outcome
	})
}

func asBinary(valA, valB int, f func(binA, binB string) string) int {
	addZeros := func(val string) string {
		for len(val) < 3 {
			val = "0" + val
		}
		return val
	}
	binA := addZeros(strconv.FormatInt(int64(valA), 2))
	binB := addZeros(strconv.FormatInt(int64(valB), 2))
	outcome := f(binA, binB)
	return int(c.Must(strconv.ParseInt(outcome, 2, 64)))
}

func getOperation(opcode int) func(int, int, *int, *Registers) []int {
	switch opcode {
	case 0: // adv
		return func(literal, combo int, instr *int, reg *Registers) []int {
			reg.a = int(math.Trunc(float64(reg.a) / math.Pow(2.0, float64(combo))))
			fmt.Printf(" => A = %v\n", reg.a)
			return nil
		}
	case 1: // bxl
		return func(literal, combo int, instr *int, reg *Registers) []int {
			reg.b = bitXor(reg.b, literal)
			fmt.Printf(" => B = %v\n", reg.b)
			return nil
		}
	case 2: // bst
		return func(literal, combo int, instr *int, reg *Registers) []int {
			reg.b = combo % 8
			fmt.Printf(" => B = %v\n", reg.b)
			return nil
		}
	case 3: // jnz
		return func(literal, combo int, instr *int, reg *Registers) []int {
			if reg.a != 0 {
				*instr = literal
			}
			fmt.Printf(" => *instr = %v\n", literal)
			return nil
		}
	case 4: // bxc
		return func(literal, combo int, instr *int, reg *Registers) []int {
			reg.b = bitXor(reg.b, reg.c)
			fmt.Printf(" => B = %v\n", reg.b)
			return nil
		}
	case 5: // out
		return func(literal, combo int, instr *int, reg *Registers) []int {
			fmt.Printf(" => out = %v\n", combo%8)
			return []int{combo % 8}
		}
	case 6: // bdv
		return func(literal, combo int, instr *int, reg *Registers) []int {
			reg.b = int(math.Trunc(float64(reg.a) / math.Pow(2.0, float64(combo))))
			fmt.Printf(" => B = %v\n", reg.b)
			return nil
		}
	case 7: // cdv
		return func(literal, combo int, instr *int, reg *Registers) []int {
			reg.c = int(math.Trunc(float64(reg.a) / math.Pow(2.0, float64(combo))))
			fmt.Printf(" => C = %v\n", reg.c)
			return nil
		}

	default:
		panic("reached default branch")
	}
}

func getCombo(literal int, reg Registers) int {
	if literal < 4 || literal == 7 {
		return literal
	}
	switch literal {
	case 4:
		return reg.a
	case 5:
		return reg.a
	case 6:
		return reg.c
	default:
		panic("default branch reached")
	}
}

func run(inputfile string) {
	y := 0
	registers := Registers{}
	var instructions []int
	var pointer int
	c.ReadLines(inputfile, func(line string) {
		switch y {
		case 0:
			registers.a = parseRegister(line)
		case 1:
			registers.b = parseRegister(line)
		case 2:
			registers.c = parseRegister(line)
		case 4:
			instructions = parseInstructions(line)
		}
		y++
	})

	output := []int{}
	for pointer < len(instructions)-1 {
		startPointer := pointer
		opcode := instructions[pointer]
		literal := instructions[pointer+1]
		combo := getCombo(literal, registers)

		operation := getOperation(opcode)
		fmt.Printf("*%v => op: %v, lit: %v, combo: %v, A: %v, B: %v, C: %v", pointer, opcode, literal, combo, registers.a, registers.b, registers.c)
		opout := operation(literal, combo, &pointer, &registers)
		if pointer == startPointer {
			pointer += 2
		}
		if opout != nil {
			output = append(output, opout...)
		}
	}

	outstr := ""
	for _, o := range output {
		outstr += strconv.Itoa(o)
	}
	fmt.Printf("Program output concatenated is: %s\n", outstr)
}
