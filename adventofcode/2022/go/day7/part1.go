package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode2022/helper"
)

type dir struct {
	name    string
	parent  *dir
	subdirs map[string]*dir
	files   map[string]int
	size    int
}

func makeDir(name string, parent *dir) *dir {
	return &dir{
		name:    name,
		parent:  parent,
		subdirs: map[string]*dir{},
		files:   map[string]int{},
	}
}

func ls(curDir *dir, console []string, position *int) {
	op := console[*position]
	for *position < len(console) && !strings.HasPrefix(console[*position], "$") {
		fmt.Println("Listing directory content", console[*position])
		op = console[*position]
		if strings.HasPrefix(op, "dir") {
			dirName := strings.Split(op, " ")[1]
			curDir.subdirs[dirName] = makeDir(dirName, curDir)
		} else {
			fileData := strings.Split(op, " ")
			fileName := fileData[1]
			fileSize, _ := strconv.Atoi(fileData[0])
			curDir.files[fileName] = fileSize
		}
		*position++
	}
}

func cd(curDir *dir, op string) *dir {
	fmt.Println(op)
	dirName := strings.Split(op, " ")[2]
	if dirName == ".." {
		return curDir.parent
	} else {
		return curDir.subdirs[dirName]
	}
}

func indent(depth int) string {
	indentation := ""
	for i := 0; i < depth; i++ {
		indentation += "  "
	}
	return indentation
}

func printFs(curDir *dir, depth int) {
	fmt.Println(indent(depth), curDir.name, "(dir, size="+strconv.Itoa(curDir.size)+")")
	for name, size := range curDir.files {
		fmt.Println(indent(depth+1), name, "(file, size="+strconv.Itoa(size)+")")
	}
	for _, dir := range curDir.subdirs {
		printFs(dir, depth+1)
	}
}

func calculateDirSize(curDir *dir) int {
	sum := 0
	for _, size := range curDir.files {
		sum += size
	}
	for _, subdir := range curDir.subdirs {
		sum += calculateDirSize(subdir)
	}
	curDir.size = sum
	return sum
}

func sumTargetDirs(curDir *dir, sum int) int {
	if curDir.size < 100000 {
		sum += curDir.size
	}
	for _, subdir := range curDir.subdirs {
		sum = sumTargetDirs(subdir, sum)
	}
	return sum
}

func main() {
	console := helper.ReadLines("input.txt")

	root := makeDir("/", nil)
	curDir := root
	position := 1
	for position < len(console) {
		op := console[position]
		if strings.HasPrefix(op, "$ cd") {
			curDir = cd(curDir, op)
			position++
		}
		if strings.HasPrefix(op, "$ ls") {
			fmt.Println(op)
			position++
			ls(curDir, console, &position)
		}
	}
	calculateDirSize(root)
	printFs(root, 0)
	sumDirs := sumTargetDirs(root, 0)
	fmt.Println(sumDirs)
}
