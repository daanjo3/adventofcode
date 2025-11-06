package lib

import (
	"bufio"
	"os"
)

func ReadLine(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func ReadLines(path string, callback func(string, int)) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		callback(scanner.Text(), line)
		line++
	}
}
