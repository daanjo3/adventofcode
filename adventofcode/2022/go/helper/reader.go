package helper

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) []string {
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
