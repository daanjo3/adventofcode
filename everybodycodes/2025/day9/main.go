package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day9", FindChild, FindChildren)
}

type DNA struct {
	id       string
	sequence string
}

type Pair struct {
	child   int
	parents []int
}

func NewPair(child string, p1 string, p2 string) Pair {
	parents := []int{
		lib.Must(strconv.Atoi(p1)),
		lib.Must(strconv.Atoi(p2)),
	}
	slices.Sort(parents)
	return Pair{
		child:   lib.Must(strconv.Atoi(child)),
		parents: parents,
	}
}

func (p Pair) Key() string {
	return fmt.Sprintf("%d%d%d", p.child, p.parents[0], p.parents[1])
}

func (child DNA) isChildOf(parents []DNA) (bool, int) {
	if len(parents) != 2 {
		panic("Can't compare against not-exactly 2 parents")
	}
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
			return false, 0
		}
	}
	similarity := 1
	for _, v := range countMap {
		similarity *= v
	}
	return true, similarity
}

func (dna DNA) String() string {
	return fmt.Sprintf("%s:%s", dna.id, dna.sequence)
}

func parseSequences(inputfile string) []DNA {
	sequences := []DNA{}
	lib.ReadLines(inputfile, func(line string, index int) {
		parts := strings.Split(line, ":")
		sequences = append(sequences, DNA{id: parts[0], sequence: parts[1]})
	})
	return sequences
}

func FindChild(inputfile string) {
	sequences := parseSequences(inputfile)
	for i, child := range sequences {
		seqCp := make([]DNA, len(sequences))
		copy(seqCp, sequences)
		parents := slices.Delete(seqCp, i, i+1)

		found, score := child.isChildOf(parents)
		if !found {
			continue
		}
		fmt.Printf("Similarity score of %s with parent %s and %s: %d\n", child.id, parents[0].id, parents[1].id, score)
	}
}

func FindChildren(inputfile string) {

	sequences := parseSequences(inputfile)
	totalScore := 0
	for _, child := range sequences {
		fmt.Println("child", child.id)
		for p1Index, p1 := range sequences {
			if p1.id == child.id {
				continue
			}
			for _, p2 := range sequences[p1Index:] {
				if p2.id == child.id || p2.id == p1.id {
					continue
				}
				if ok, score := child.isChildOf([]DNA{p1, p2}); ok {
					fmt.Printf("%s is a child of %s and %s\n", child.id, p1.id, p2.id)
					totalScore += score
				}
			}
		}
	}
	fmt.Printf("Total score of complete tree is %d\n", totalScore)
}
