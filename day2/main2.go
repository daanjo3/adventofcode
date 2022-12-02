package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type Gesture int64

const (
	Rock    Gesture = 0
	Paper           = 1
	Scissor         = 2
)

func parseGesture(key string) Gesture {
	switch key {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissor
	}
	panic("This should never be reached")
}

// X: i lose
// Y: draw
// Z: i win

func getMyGesture(opponent Gesture, outcome string) Gesture {
	switch outcome {
	case "X":
		if opponent == 0 {
			return opponent + 2
		}
		return opponent - 1
	case "Y":
		return opponent
	case "Z":
		if opponent == 2 {
			return opponent - 2
		}
		return opponent + 1
	}
	panic("This should never be reached")
}

func getMatchPoint(outcome string) int {
	switch outcome {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	}
	panic("This should never be reached")
}

// Can't believe this one is necessary
func getGesturePoint(gesture Gesture) int {
	switch gesture {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissor:
		return 3
	}
	panic("This should never be reached")
}

func main() {
	matches := readFile("input.txt")

	total := 0
	for _, match := range matches {
		codes := strings.Fields(match)
		opponent := parseGesture(codes[0])
		me := getMyGesture(opponent, codes[1])
		matchPoint := getMatchPoint(codes[1])
		points := matchPoint + getGesturePoint(me)
		fmt.Println(codes, points)
		total += points
	}
	fmt.Println(total)
}
