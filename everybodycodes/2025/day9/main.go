package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day9", FindChild)
}

type DNA struct {
	id       string
	sequence string
}

func (child DNA) CheckParents(parents []DNA) (map[string]int, error) {
	countMap := map[string]int{}

	for i, symbol := range child.sequence {
		found := false
		for _, parent := range parents {
			if rune(parent.sequence[i]) == symbol {
				found = true
				countMap[parent.id]++
			}
		}
		if !found {
			return nil, fmt.Errorf("sequences %s does not have a match with %v at position %d", child, maps.Keys(countMap), i)
		}
	}
	return countMap, nil
}

func (dna DNA) String() string {
	return fmt.Sprintf("%s:%s", dna.id, dna.sequence)
}

func FindChild(inputfile string) {
	sequences := []DNA{}
	lib.ReadLines(inputfile, func(line string, index int) {
		parts := strings.Split(line, ":")
		sequences = append(sequences, DNA{id: parts[0], sequence: parts[1]})
	})
	for i, child := range sequences {
		seqCp := make([]DNA, len(sequences))
		copy(seqCp, sequences)
		parents := slices.Delete(seqCp, i, i+1)

		countMap, err := child.CheckParents(parents)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Child %s matches parents: %v\n", child.id, countMap)
		score := 1
		for _, count := range countMap {
			score *= count
		}
		fmt.Printf("Similarity score: %d\n", score)
	}
}
