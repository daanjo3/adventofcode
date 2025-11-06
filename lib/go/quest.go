package lib

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Quest(dayDirName string, partFuncs ...func(string)) {
	var part int
	var sample bool

	inputMsgPart := ""
	partMsgPart := ""
	flag.IntVar(&part, "part", 1, "Problem part to run")
	flag.BoolVar(&sample, "sample", false, "Run with sample input")
	flag.Parse()

	if part < 1 || part > len(partFuncs) {
		fmt.Printf("-part must be between 1 and %d\n", len(partFuncs))
		os.Exit(1)
	}
	partFunc := partFuncs[part-1]
	partMsgPart = strconv.Itoa(part)
	if sample {
		inputMsgPart = "sample"
	} else {
		inputMsgPart = "full"
	}

	inputPath := getInputPath(dayDirName, part, sample)

	fmt.Printf("Running %s: part=%s mode=%s input=%s\n", dayDirName, partMsgPart, inputMsgPart, inputPath)
	start := time.Now()
	defer func() {
		fmt.Printf("Took %s\n", time.Since(start))
	}()
	partFunc(inputPath)
}

var (
	inputPathFormat           = "%s/input.txt"
	inputPartPathFormat       = "%s/input-%d.txt"
	sampleInputPathFormat     = "%s/input-sample.txt"
	sampleInputPartPathFormat = "%s/input-sample-%d.txt"
)

func getInputPath(dayDirName string, part int, sample bool) string {
	if sample {
		sampleInputPartFp := fmt.Sprintf(sampleInputPartPathFormat, dayDirName, part)
		if _, err := os.Stat(sampleInputPartFp); err == nil {
			return sampleInputPartFp
		}
		return fmt.Sprintf(sampleInputPathFormat, dayDirName)
	}
	partInputFp := fmt.Sprintf(inputPartPathFormat, dayDirName, part)
	if _, err := os.Stat(partInputFp); err == nil {
		return partInputFp
	}
	return fmt.Sprintf(inputPathFormat, dayDirName)
}

func PlaceholderFunc(inputfile string) {
	fmt.Println("To be implemented...")
}
