package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day9", FindChild, FindChildren, IndexFamily)
}

type DragonDuck struct {
	scaleNumber int
	dna         string
	family      *string
	parents     []*DragonDuck
	children    []*DragonDuck
}

func (dd *DragonDuck) isChildOf(parents []*DragonDuck) (bool, int) {
	if len(parents) != 2 {
		panic("Can't compare against not-exactly 2 parents")
	}
	countMap := map[int]int{}
	for i, symbol := range dd.dna {
		found := false
		for _, parent := range parents {
			if rune(parent.dna[i]) == symbol {
				found = true
				countMap[parent.scaleNumber]++
			}
		}
		if !found {
			return false, 0
		}
	}
	for _, parent := range parents {
		parent.AddChild(dd)
	}
	dd.AddParents(parents)
	similarity := 1
	for _, v := range countMap {
		similarity *= v
	}
	return true, similarity
}

func (dd *DragonDuck) AddParents(parents []*DragonDuck) {
	dd.parents = parents
}

func (dd *DragonDuck) AddChild(child *DragonDuck) {
	dd.children = append(dd.children, child)
}

func (dd *DragonDuck) IndexFamily(family string) {
	if dd.family != nil {
		return
	}
	dd.family = &family
	for _, parent := range dd.parents {
		parent.IndexFamily(family)
	}
	for _, child := range dd.children {
		child.IndexFamily(family)
	}
}

func (dd DragonDuck) String() string {
	return fmt.Sprintf("%s:%s of the %s", dd.scaleNumber, dd.dna, *dd.family)
}

func parseSequences(inputfile string) []*DragonDuck {
	sequences := []*DragonDuck{}
	lib.ReadLines(inputfile, func(line string, index int) {
		parts := strings.Split(line, ":")
		sequences = append(sequences, &DragonDuck{scaleNumber: lib.Must(strconv.Atoi(parts[0])), dna: parts[1]})
	})
	return sequences
}

func FindChild(inputfile string) {
	dragonducks := parseSequences(inputfile)
	for i, child := range dragonducks {
		seqCp := make([]*DragonDuck, len(dragonducks))
		copy(seqCp, dragonducks)
		parents := slices.Delete(seqCp, i, i+1)

		found, score := child.isChildOf(parents)
		if !found {
			continue
		}
		fmt.Printf("Similarity score of %s with parent %s and %s: %d\n", child.scaleNumber, parents[0].scaleNumber, parents[1].scaleNumber, score)
	}
}

func FindChildren(inputfile string) {

	dragonducks := parseSequences(inputfile)
	totalScore := 0
	for _, child := range dragonducks {
		fmt.Println("child", child.scaleNumber)
		for p1Index, p1 := range dragonducks {
			if p1.scaleNumber == child.scaleNumber {
				continue
			}
			for _, p2 := range dragonducks[p1Index:] {
				if p2.scaleNumber == child.scaleNumber || p2.scaleNumber == p1.scaleNumber {
					continue
				}
				if ok, score := child.isChildOf([]*DragonDuck{p1, p2}); ok {
					fmt.Printf("%s is a child of %s and %s\n", child.scaleNumber, p1.scaleNumber, p2.scaleNumber)
					totalScore += score
				}
			}
		}
	}
	fmt.Printf("Total score of complete tree is %d\n", totalScore)
}

func IndexFamily(inputfile string) {

	dragonducks := parseSequences(inputfile)
	for _, child := range dragonducks {
		for p1Index, p1 := range dragonducks {
			if p1.scaleNumber == child.scaleNumber {
				continue
			}
			for _, p2 := range dragonducks[p1Index:] {
				if p2.scaleNumber == child.scaleNumber || p2.scaleNumber == p1.scaleNumber {
					continue
				}
				child.isChildOf([]*DragonDuck{p1, p2})
			}
		}
	}

	for i, dragonduck := range dragonducks {
		if dragonduck.family == nil {
			// TODO implement joining families?
			dragonduck.IndexFamily(fmt.Sprintf("The %dth order", i))
		}
	}
	families := map[string][]*DragonDuck{}
	for _, dragonduck := range dragonducks {
		if _, ok := families[*dragonduck.family]; !ok {
			families[*dragonduck.family] = []*DragonDuck{}
		}
		families[*dragonduck.family] = append(families[*dragonduck.family], dragonduck)
	}
	largestFamilyName := ""
	largestFamily := 0
	for familyName, members := range families {
		if len(members) > largestFamily {
			largestFamilyName = familyName
		}
	}
	largestFamilySize := 0
	for _, member := range families[largestFamilyName] {
		largestFamilySize += member.scaleNumber
	}
	fmt.Printf("The largest family has size %d\n", largestFamilySize)
}
