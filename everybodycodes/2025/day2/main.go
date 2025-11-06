package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day2", FirstPage, lib.PlaceholderFunc, lib.PlaceholderFunc)
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

func FirstPage(filename string) {
	numbers := ReadInput(filename)
	total := DoFirstPageArithmetic(numbers)
	fmt.Printf("The first page of calculations results in %s\n", total)
}

func DoFirstPageArithmetic(a ComplexNumber) ComplexNumber {
	total := ComplexNumber{X: 0, Y: 0}
	for i := 0; i < 3; i++ {
		total = total.Multiply(total)
		total = total.Divide(ComplexNumber{X: 10, Y: 10})
		total = total.Add(a)
	}
	return total
}

func ReadInput(filename string) ComplexNumber {
	return ParseComplexNumber(lib.ReadLine(filename))
}
