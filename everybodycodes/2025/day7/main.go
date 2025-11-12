package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/daanjo3/adventofcode/lib"
)

func main() {
	lib.Quest("day7", findSingleName, findAllNames, generateNames)
}

type RuleMap map[rune][]rune

func (rm RuleMap) NameAllowed(name string) bool {
	for i, c := range name {
		if i == len(name)-1 {
			return true
		}
		if !rm.Allowed(c, rune(name[i+1])) {
			fmt.Printf("Name '%s' is not allowed, char '%c' at pos %d break rule: %s\n", name, c, i, rm.StringEntry(c))
			return false
		}
	}
	panic("This should be unreachable")
}

func (rm RuleMap) Allowed(key rune, next rune) bool {
	values := rm[key]
	return slices.Contains(values, next)
}

func (rm RuleMap) StringEntry(key rune) string {
	return fmt.Sprintf("%c > %s", key, string(rm[key]))
}

func findSingleName(inputfile string) {
	names, rules := parseNotes(inputfile)
	checkNames(names, rules)
}

func findAllNames(inputfile string) {
	names, rules := parseNotes(inputfile)
	accepted := checkNames(names, rules)
	sum := 0
	for i, name := range names {
		if slices.Contains(accepted, name) {
			sum += i + 1
		}
	}
	fmt.Printf("The sum of accepted name indexes is %d\n", sum)
}

func generateNames(inputfile string) {
	names, rules := parseNotes(inputfile)
	all := []string{}
	for _, name := range names {
		if !rules.NameAllowed(name) {
			continue
		}
		found := []string{}
		considerPrefixes([]rune(name), len(name), 1, rune(name[len(name)-1]), rules, &found)
		all = append(all, found...)
	}
	all = lib.ArrUnique(all)
	fmt.Printf("Total amount of unique names that can be created: %d\n", len(all))
}

func considerPrefixes(name []rune, length int, options int, cur rune, rules RuleMap, found *[]string) {
	fmt.Println(string(name))
	if length >= 7 {
		*found = append(*found, string(name))
	}
	letters := rules[cur]
	if len(letters) == 0 || length == 11 {
		return
	}
	for _, letter := range letters {
		newName := append(name, letter)
		considerPrefixes(newName, length+1, options, letter, rules, found)
	}
}

func checkNames(names []string, rules RuleMap) []string {
	accepted := []string{}
	for _, name := range names {
		if rules.NameAllowed(name) {
			accepted = append(accepted, name)
		}
	}
	fmt.Printf("Name %s follows the ruleset\n", accepted)
	return accepted
}

func parseNotes(inputfile string) (names []string, rules RuleMap) {
	rules = RuleMap{}
	lib.ReadLines(inputfile, func(line string, index int) {
		if index == 0 {
			names = strings.Split(line, ",")
		}
		if index > 1 {
			parts := strings.Split(line, " > ")
			source := rune(parts[0][0])
			targets := []rune{}
			for _, r := range strings.Split(parts[1], ",") {
				targets = append(targets, rune(r[0]))
			}
			rules[source] = targets
		}
	})
	return names, rules
}
