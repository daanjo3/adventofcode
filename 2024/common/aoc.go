package common

import (
	"flag"
	"fmt"
	"os"
)

func AdventCommand(dayDir string, oneStarFunc func(string), twoStarFunc func(string)) {
	var inputfile string
	var star int
	var starFunc func(string)
	var sample bool

	inputMsgPart := ""
	starMsgPart := ""
	flag.IntVar(&star, "star", 1, "Star problem to run (1 or 2)")
	flag.BoolVar(&sample, "sample", false, "Run with sample input")
	flag.Parse()

	if star != 1 && star != 2 {
		fmt.Println("-star must be either 1 or 2")
		os.Exit(1)
	}
	if star == 1 {
		starFunc = oneStarFunc
		starMsgPart = "*"
	} else {
		starFunc = twoStarFunc
		starMsgPart = "**"
	}

	if sample {
		inputfile = fmt.Sprintf("%s/input-sample.txt", dayDir)
		inputMsgPart = "sample"
	} else {
		inputfile = fmt.Sprintf("%s/input.txt", dayDir)
		inputMsgPart = "full"
	}

	fmt.Printf("Running %s: part=%s input=%s\n", dayDir, starMsgPart, inputMsgPart)
	starFunc(inputfile)
}
