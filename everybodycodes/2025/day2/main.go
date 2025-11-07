package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day2", FirstPage, EngraveBreastPlate, lib.PlaceholderFunc)
}

type ComplexNumber struct {
	X int
	Y int
}

func (a ComplexNumber) Add(b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a ComplexNumber) Multiply(b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		X: a.X*b.X - a.Y*b.Y,
		Y: a.X*b.Y + a.Y*b.X,
	}
}

func (a ComplexNumber) Divide(b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		X: (a.X / b.X),
		Y: (a.Y / b.Y),
	}
}

func (c ComplexNumber) String() string {
	return fmt.Sprintf("[%d,%d]", c.X, c.Y)
}

func ParseComplexNumber(str string) ComplexNumber {
	match := lib.Must(regexp.Compile("A=\\[(.*)\\]")).FindSubmatch([]byte(str))[1]
	numbers := strings.Split(string(match), ",")
	return ComplexNumber{
		X: lib.Must(strconv.Atoi(numbers[0])),
		Y: lib.Must(strconv.Atoi(numbers[1])),
	}
}

func ReadInput(filename string) ComplexNumber {
	return ParseComplexNumber(lib.ReadLine(filename))
}

func FirstPage(filename string) {
	numbers := ReadInput(filename)
	total := DoSampleCalculations(numbers)
	fmt.Printf("The first page of calculations results in %s\n", total)
}

func DoSampleCalculations(a ComplexNumber) ComplexNumber {
	total := ComplexNumber{X: 0, Y: 0}
	for i := 0; i < 3; i++ {
		total = total.Multiply(total)
		total = total.Divide(ComplexNumber{X: 10, Y: 10})
		total = total.Add(a)
	}
	return total
}

func EngraveBreastPlate(filename string) {
	start := ReadInput(filename)
	distance := 1000
	step := 10
	maxCycles := 100

	totalEngraved := 0
	for y := 0; y <= distance; y += step {
		for x := 0; x <= distance; x += step {
			current := start.Add(ComplexNumber{X: x, Y: y})
			doEngrave := true
			for i := 0; i < maxCycles; i++ {
				current = current.Multiply(current)
				fmt.Println("multiplied:", current)
				current = current.Divide(ComplexNumber{X: 100_000, Y: 100_000})
				fmt.Println("divided:", current)
				current = current.Add(start)
				fmt.Println("added:", current)
				if current.X < -1_000_000 || current.X > 1_000_000 || current.Y < -1_000_000 || current.Y > 1_000_000 {
					doEngrave = false
					fmt.Printf("P=%s R=%s C=%d\n", start.Add(ComplexNumber{X: x, Y: y}), current, i)
					break
				}
				return
			}
			if doEngrave {
				fmt.Printf("P=%s R=%s\n", start.Add(ComplexNumber{X: x, Y: y}), current)
				totalEngraved++
			}
		}
		// fmt.Println()
	}
	fmt.Println("Total engraved points:", totalEngraved)
}
