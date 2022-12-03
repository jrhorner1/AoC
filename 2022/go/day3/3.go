package day3

import (
	"strings"

	"github.com/jrhorner1/AoC/pkg/sets"
	"github.com/sirupsen/logrus"
)

type Rucksack struct {
	allItems,
	compartment1,
	compartment2 sets.Rune
}

type Rucksacks []Rucksack

func Puzzle(input *[]byte, part2 bool) int {
	var rucksacks Rucksacks
	for _, items := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		var rucksack Rucksack
		rucksack.allItems = sets.NewRune()
		rucksack.compartment1 = sets.NewRune()
		rucksack.compartment2 = sets.NewRune()
		for _, item := range items {
			rucksack.allItems.Insert(item)
		}
		for _, item := range items[:(len(items) / 2)] {
			rucksack.compartment1.Insert(item)
		}
		for _, item := range items[(len(items) / 2):] {
			rucksack.compartment2.Insert(item)
		}
		rucksacks = append(rucksacks, rucksack)
	}
	logrus.Debug(rucksacks)
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
		itemSet := rucksack.compartment1.Intersection(rucksack.compartment2)
		if itemSet.Len() != 1 {
			logrus.Fatal("Something's wrong here, I can feel it!")
		}
		item, _ := itemSet.PopAny()
		score += priority(item)
	}
	return score
}

func (rucksacks *Rucksacks) findBadges() int {
	score := 0
	for i := 0; i < len(*rucksacks); i += 3 {
		itemSet := (*rucksacks)[i].allItems.Intersection(
			(*rucksacks)[i+1].allItems.Intersection((*rucksacks)[i+2].allItems),
		)
		if itemSet.Len() != 1 {
			logrus.Fatal("Something's wrong here, I can feel it!")
		}
		item, _ := itemSet.PopAny()
		score += priority(item)
	}
	return score
}
