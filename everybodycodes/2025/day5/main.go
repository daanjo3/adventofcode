package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day5", AssessArmourersSword, CompareSwords, CompareSwordsExtreme)
}

type Level struct {
	Core int
	Low  int
	High int
}

func (l *Level) Add(part int) bool {
	if part < l.Core && l.Low == 0 {
		l.Low = part
		return true
	}
	if part > l.Core && l.High == 0 {
		l.High = part
		return true
	}
	return false
}

func (l Level) String() string {
	str := ""
	if l.Low > 0 {
		str += strconv.Itoa(l.Low) + "-"
	} else {
		str += "  "
	}
	str += strconv.Itoa(l.Core)
	if l.High > 0 {
		str += "-" + strconv.Itoa(l.High)
	}
	return str
}

func (l Level) Value() int {
	str := ""
	if l.Low > 0 {
		str += strconv.Itoa(l.Low)
	}
	str += strconv.Itoa(l.Core)
	if l.High > 0 {
		str += strconv.Itoa(l.High)
	}
	return lib.Must(strconv.Atoi(str))
}

func (l Level) Compare(l2 Level) int {
	return l.Value() - l2.Value()
}

type Sword struct {
	ID     int
	Levels []Level
}

func (s *Sword) Add(part int) {
	for i := range s.Levels {
		if s.Levels[i].Add(part) {
			return
		}
	}
	s.Levels = append(s.Levels, Level{Core: part})
}

func (s Sword) String() string {
	str := ""
	for i, level := range s.Levels {
		str += fmt.Sprintf("%s\n", level.String())
		if i < len(s.Levels)-1 {
			str += "  |\n"
		}
	}
	return str
}

func (s Sword) Quality() int {
	quality := ""
	for _, level := range s.Levels {
		quality += strconv.Itoa(level.Core)
	}
	return lib.Must(strconv.Atoi(quality))
}

func (s Sword) Compare(s2 Sword) int {
	diff := s.Quality() - s2.Quality()
	if diff != 0 {
		return diff
	}
	for i := 0; i < len(s.Levels) && i < len(s2.Levels); i++ {
		levelDiff := s.Levels[i].Compare(s2.Levels[i])
		if levelDiff != 0 {
			return levelDiff
		}
	}
	return s.ID - s2.ID
}

func NewSword(id int, parts []int) Sword {
	sword := Sword{
		ID: id,
	}
	for _, part := range parts {
		sword.Add(part)
	}
	return sword
}

func ParseSword(line string) Sword {
	idAndParts := strings.Split(line, ":")
	id := lib.Must(strconv.Atoi(idAndParts[0]))

	parts := []int{}
	for _, part := range strings.Split(idAndParts[1], ",") {
		parts = append(parts, lib.Must(strconv.Atoi(part)))
	}
	return NewSword(id, parts)
}

func AssessArmourersSword(inputfile string) {
	line := lib.ReadLine(inputfile)
	sword := ParseSword(line)
	fmt.Printf("Armourer's Sword Structure:\n%s\n", sword.String())
	fmt.Printf("Armourer's Sword Quality: %d\n", sword.Quality())
}

func CompareSwords(inputfile string) {
	swords := []Sword{}
	lib.ReadLines(inputfile, func(line string, _ int) {
		swords = append(swords, ParseSword(line))
	})
	slices.SortFunc(swords, func(a, b Sword) int {
		return a.Quality() - b.Quality()
	})
	fmt.Printf("Best sword: %d\n", swords[len(swords)-1].Quality())
	fmt.Printf("Worst sword: %d\n", swords[0].Quality())
	fmt.Printf("Quality difference: %d\n", swords[len(swords)-1].Quality()-swords[0].Quality())
}

func CompareSwordsExtreme(inputfile string) {
	swords := []Sword{}
	lib.ReadLines(inputfile, func(line string, _ int) {
		swords = append(swords, ParseSword(line))
	})
	slices.SortFunc(swords, func(a, b Sword) int {
		return b.Compare(a)
	})
	fmt.Printf("Worst sword: %d\n", swords[len(swords)-1].Quality())
	fmt.Printf("Best sword: %d\n", swords[0].Quality())
	checksum := 0
	for i, sword := range swords {
		checksum += sword.ID * (i + 1)
	}
	fmt.Printf("Checksum: %d\n", checksum)
}
