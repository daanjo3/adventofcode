package main

import (
	"fmt"
	"strconv"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day9",
		compressFileSystem,
		c.PlaceholderFunc,
	)
}

func parseFs(config string) ([]int, int) {
	fs := []int{}
	isFile := true
	fileNum := 0
	for _, char := range config {
		digit := c.Must(strconv.Atoi(string(char)))
		for range digit {
			if isFile {
				fs = append(fs, fileNum)
			} else {
				fs = append(fs, -1)
			}
		}
		if isFile {
			isFile = false
			fileNum++
		} else {
			isFile = true
		}
	}
	return fs
}

func canRelocate(fs []int) bool {
	foundFree := false
	for _, v := range fs {
		if foundFree && v != -1 {
			return true
		}
		if v == -1 {
			foundFree = true
		}
	}
	return false
}

func nextFreeIndex(fs []int) int {
	for i, v := range fs {
		if v == -1 {
			return i
		}
	}
	panic("Didn't find free blocks")
}

func nextFreeBlockIndex(fs []int, blockSize int) int {
	for i, v := range fs {
		freeSize := 0
		if v == -1 {
			for end := i; fs[end] == -1; end++ {
				freeSize++
				if freeSize == blockSize {
					return i
				}
			}
		}
	}
	return -1
}

func popLastFile(fs []int) int {
	for i := len(fs) - 1; i >= 0; i-- {
		v := fs[i]
		if v != -1 {
			fs[i] = -1
			return v
		}
	}
	panic("Didn't find file blocks")
}

func getFileBlock(fs []int, blockNum int) (int, int) {
	start, stop := -1
	for i, v := range fs {
		if start == -1 && v == blockNum {
			start = i
		}
		if start != -1 && v != blockNum {
			stop = i - 1
			return start, stop
		}
	}
	panic("Didn't find file block: " + strconv.Itoa(blockNum))
}

func calculateChecksum(fs []int) int {
	sum := 0
	for i, v := range fs {
		if v == -1 {
			continue
		}
		sum += i * v
	}
	return sum
}

func compressFileSystem(inputfile string) {
	line := c.ReadLine(inputfile)
	fs, _ := parseFs(line)
	// fmt.Printf("config filesystem: %v\n", line)
	// fmt.Printf("parsed filesystem: %v\n", fs)

	for canRelocate(fs) {
		iNextFree := nextFreeIndex(fs)
		toRelocate := popLastFile(fs)
		fs[iNextFree] = toRelocate
	}

	// fmt.Println("compressed filesystem", fs)
	fmt.Printf("The compressed filesystem checksum is %v.\n", calculateChecksum(fs))
}

func compressFileSystemBlocks(inputfile string) {
	line := c.ReadLine(inputfile)
	fs, blockNums := parseFs(line)
	// fmt.Printf("config filesystem: %v\n", line)
	// fmt.Printf("parsed filesystem: %v\n", fs)

	for canRelocate(fs) {
		toRelocate := getFileBlock(fs)
		iNextFree := nextFreeBlockIndex(fs, len(toRelocate))
		// if iNextFree != -1 {
		// 	for
		// }

		fs[iNextFree] = toRelocate
	}

	// fmt.Println("compressed filesystem", fs)
	fmt.Printf("The compressed filesystem checksum is %v.\n", calculateChecksum(fs))
}
