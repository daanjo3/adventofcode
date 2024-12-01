package common

import (
	"bufio"
	"os"
)

func ReadLines(path string, callback func(string)) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}
}
