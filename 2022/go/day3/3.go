package day3

import (
	"strings"
)

type Rucksack struct {
	allItems,
	compartment1,
	compartment2 string
}

func Puzzle(input *[]byte, part2 bool) int {
	score := 0
	var rucksacks []Rucksack
	for _, items := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		var rucksack Rucksack
		rucksack.allItems = items
		rucksack.compartment1 = items[:(len(items) / 2)]
		rucksack.compartment2 = items[(len(items) / 2):]
		rucksacks = append(rucksacks, rucksack)
	}
	if part2 {
	findBadges:
		for i := 0; i < len(rucksacks); i += 3 {
			for _, item1 := range rucksacks[i].allItems {
				for _, item2 := range rucksacks[i+1].allItems {
					for _, item3 := range rucksacks[i+2].allItems {
						if item1 == item2 && item2 == item3 {
							score += priority(item3)
							continue findBadges
						}
					}
				}
			}
		}
		return score
	}
	for _, rucksack := range rucksacks {
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
