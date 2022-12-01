package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readFile(filePath string) []string {
	// Open the file
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set the scanner to split on lines stripping trailing whitespaces
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	// Actually go through the file by calling Scan() until it is false
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	// Close the file
	readFile.Close()

	return fileLines
}

func calculateElfCapacities(calories []string) []int {
	elfCapacity := []int{0}
	elfSum := 0
	for _, line := range calories {
		if line == "" {
			elfCapacity = append(elfCapacity, elfSum)
			elfSum = 0
		} else {
			parsed, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			elfSum += parsed
		}
	}
	return elfCapacity
}

func main() {
	calories := readFile("input.txt")
	elfCapacity := calculateElfCapacities(calories)

	// Reverse sort the list
	sort.Sort(sort.Reverse(sort.IntSlice(elfCapacity)))

	// Sum the top 3
	total := elfCapacity[0] + elfCapacity[1] + elfCapacity[2]
	fmt.Println(total)
}
