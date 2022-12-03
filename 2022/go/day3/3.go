package day3

import (
	"regexp"
	"strings"
)

type Rucksack struct {
	allItems,
	compartment1,
	compartment2 string
}

type Rucksacks []Rucksack

func Puzzle(input *[]byte, part2 bool) int {
	var rucksacks Rucksacks
	for _, items := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		var rucksack Rucksack
		rucksack.allItems = items
		rucksack.compartment1 = items[:(len(items) / 2)]
		rucksack.compartment2 = items[(len(items) / 2):]
		rucksacks = append(rucksacks, rucksack)
	}
	if part2 {
		return rucksacks.findBadges()
	}
	return rucksacks.findItem()
}

func priority(item rune) int {
	lowerCasePriority := int('a') - 1
	upperCasePriority := int('A') - 27
	ascii := int(item)
	if ascii >= int('a') && ascii <= int('z') {
		return ascii - lowerCasePriority
	} else if ascii >= int('A') && ascii <= int('Z') {
		return ascii - upperCasePriority
	}
	return 0
}

func (rucksacks *Rucksacks) findItem() int {
	score := 0
	for _, rucksack := range *rucksacks {
	compare:
		for _, item1 := range rucksack.compartment1 {
			for _, item2 := range rucksack.compartment2 {
				if item2 == item1 {
					score += priority(item2)
					break compare
				}
			}
		}
	}

	return score
}

func (rucksacks *Rucksacks) findBadges() int {
	score := 0
groups:
	for i := 0; i < len(*rucksacks); i += 3 {
		for _, item1 := range (*rucksacks)[i].allItems {
			re := regexp.MustCompile(string(item1))
			if re.MatchString((*rucksacks)[i+1].allItems) && re.MatchString((*rucksacks)[i+2].allItems) {
				score += priority(item1)
				continue groups
			}
		}
	}
	return score
}
