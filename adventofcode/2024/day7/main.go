package main

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day7", calculateCalibrationResult, calculateCalibrationResultConcat)
}

func parseLine(line string) (int, []int) {
	parts := strings.Split(line, ": ")
	total := c.Must(strconv.Atoi(parts[0]))
	values := []int{}
	for _, val := range strings.Fields(parts[1]) {
		values = append(values, c.Must(strconv.Atoi((val))))
	}
	return total, values
}

func concatInt(numA int, numB int) int {
	total := strconv.Itoa(numA) + strconv.Itoa(numB)
	return c.Must(strconv.Atoi(total))
}

func calculateWithConcat(target int, total int, values []int) bool {
	if len(values) == 1 {
		return target == total+values[0] || target == total*values[0] || target == concatInt(total, values[0])
	}

	return calculateWithConcat(target, total+values[0], values[1:]) ||
		calculateWithConcat(target, total*values[0], values[1:]) ||
		calculateWithConcat(target, concatInt(total, values[0]), values[1:])
}

func calculate(target int, total int, values []int) bool {
	if len(values) == 1 {
		return target == total+values[0] || target == total*values[0]
	}
	return calculate(target, total+values[0], values[1:]) || calculate(target, total*values[0], values[1:])
}

func _calculateCalibrationResult(inputfile string, calculateFunc func(int, int, []int) bool) {
	sum := 0

	c.ReadLines(inputfile, func(line string) {
		target, values := parseLine(line)
		fmt.Printf("%s => \n", line)
		if calculateFunc(target, values[0], values[1:]) {
			sum += target
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	})

	fmt.Printf("The sum of correctly calibrated values is %v.\n", sum)
}

func calculateCalibrationResultConcat(inputfile string) {
	_calculateCalibrationResult(inputfile, calculateWithConcat)
}

func calculateCalibrationResult(inputfile string) {
	_calculateCalibrationResult(inputfile, calculate)
}
