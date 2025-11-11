package main

import (
	"fmt"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day6", FindSwordfighterMentors, FindAllMentors, FindCloseMentors)
}

type Troop struct {
	garrison *[]Troop
	class    rune
	position int
}

func (t Troop) FindMentors(novice rune, maxDistance int) int {
	count := 0
	if t.IsMentorOf(novice) {
		count = 1
	}
	if t.position == 0 || maxDistance == 0 {
		return count
	}
	return count + (*t.garrison)[t.position-1].FindMentors(novice, maxDistance-1)
}

func (t Troop) IsMentorOf(novice rune) bool {
	return t.class == novice-32
}

func (t Troop) IsNovice() bool {
	return t.class >= 'a' && t.class <= 'z'
}

func (t Troop) GarrisonString() string {
	var result string
	for i, troop := range *t.garrison {
		if i == t.position {
			result += fmt.Sprintf("(%c)", troop.class)
		} else {
			result += fmt.Sprintf("%c", troop.class)
		}
	}
	return result
}

func FindSwordfighterMentors(inputfile string) {
	garrison := ParseGarrison(inputfile)
	totalMentorCombinations := 0
	for _, troop := range garrison {
		if troop.class != 'a' {
			continue
		}
		mentors := troop.FindMentors(troop.class, 0)
		totalMentorCombinations += mentors
		fmt.Printf("%s -> mentors: %d\n", troop.GarrisonString(), mentors)
	}
	fmt.Println("Total mentor combinations:", totalMentorCombinations)
}

func FindAllMentors(inputfile string) {
	garrison := ParseGarrison(inputfile)
	totalMentorCombinations := 0
	for _, troop := range garrison {
		if !troop.IsNovice() {
			continue
		}
		mentors := troop.FindMentors(troop.class, 0)
		totalMentorCombinations += mentors
		fmt.Printf("%s -> mentors: %d\n", troop.GarrisonString(), mentors)
	}
	fmt.Println("Total mentor combinations:", totalMentorCombinations)
}

func FindCloseMentors(inputfile string) {
	sizeFactor := 1000
	maxDistance := 1000
	if strings.Contains(inputfile, "sample") {
		sizeFactor = 3
		maxDistance = 10
	}

	garrison := ParseGarrisonElongated(inputfile, sizeFactor)
	totalMentorCombinations := 0
	for _, troop := range garrison {
		if !troop.IsNovice() {
			continue
		}
		mentors := troop.FindMentors(troop.class, maxDistance)
		totalMentorCombinations += mentors
		fmt.Printf("%s -> mentors: %d\n", troop.GarrisonString(), mentors)
	}
	fmt.Println("Total mentor combinations:", totalMentorCombinations)
}

func ParseGarrison(inputfile string) []Troop {
	garrison := []Troop{}
	line := lib.ReadLine(inputfile)
	for i, char := range line {
		garrison = append(garrison, Troop{class: char, position: i, garrison: &garrison})
	}
	return garrison
}

func ParseGarrisonElongated(inputfile string, times int) []Troop {
	garrison := []Troop{}
	line := lib.ReadLine(inputfile)
	for i := 0; i < times; i++ {
		for j, char := range line {
			garrison = append(garrison, Troop{class: char, position: (i * len(line)) + j, garrison: &garrison})
		}
	}

	return garrison
}
