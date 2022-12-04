package day4

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Elves []Elf

type Elf struct {
	sections sets.Int
}

func Puzzle(input *[]byte, part2 bool) int {
	fullyContained := 0
	overlap := 0
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		sectionRanges := strings.Split(strings.TrimSpace(line), ",")
		elves := Elves{}
		for _, sectionRange := range sectionRanges {
			sections_str := strings.Split(strings.TrimSpace(sectionRange), "-")
			sectionSet := sets.NewInt()
			sections := []int{}
			for _, section_str := range sections_str {
				section, err := strconv.Atoi(section_str)
				if err != nil {
					logrus.Error(err)
				}
				sectionSet.Insert(section)
				sections = append(sections, section)
			}
			for i := sections[0] + 1; i < sections[1]; i++ {
				sectionSet.Insert(i)
			}
			elf := Elf{sectionSet}
			elves = append(elves, elf)
		}
		if elves[0].sections.IsSuperset(elves[1].sections) || elves[1].sections.IsSuperset(elves[0].sections) {
			fullyContained++
		}
		if elves[0].sections.Intersection(elves[1].sections).Len() > 0 || elves[1].sections.Intersection(elves[0].sections).Len() > 0 {
			overlap++
		}
	}
	if part2 {
		return overlap
	}
	return fullyContained
}
