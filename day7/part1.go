package main

import (
	"strings"

	"github.com/daanjo3/adventofcode2022/helper"
)

type folder struct {
	depth  int
	parent *map[string]folder
	size   int
}

func makeFolder() *folder {
	parentFolder := make(map[string]folder)
	return &folder{
		depth:  0,
		parent: &parentFolder,
		size:   0,
	}
}

func cd(curDir *folder, dirName string) *folder {
	if (dirName)
}

func main() {
	console := helper.ReadLines("input-short.txt")
	curDir := makeFolder()

	position := 1
	for position < len(console) {
		op := console[position]
		if strings.HasPrefix(op, "$ cd") {
			dirName := strings.Split(op, " ")[2]
			curDir = cd(curDir, dirName)
			position++
		}
		if strings.HasPrefix(op, "$ ls") {

		}
	}
}
