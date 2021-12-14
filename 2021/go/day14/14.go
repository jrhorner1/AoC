package day14

import (
	"sort"
	"strings"
)

func Puzzle(input *[]byte, rounds int) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n\n")
	pairs := make(map[string]int)
	for i := 1; i < len(in[0]); i++ {
		pair := string([]byte{in[0][i-1], in[0][i]})
		pairs[pair] += 1
	}
	rules := make(map[string]string)
	for _, line := range strings.Split(strings.TrimSpace(in[1]), "\n") {
		l := strings.Split(line, " -> ")
		rules[l[0]] = l[1]
	}
	for round := rounds; round > 0; round-- {
		newPairs := make(map[string]int)
		for pair, count := range pairs {
			if insert, found := rules[pair]; found {
				newPairs[string([]byte{pair[0], insert[0]})] += count
				newPairs[string([]byte{insert[0], pair[1]})] += count
			} else {
				newPairs[pair] = count
			}
		}
		pairs = newPairs
	}
	elements := make(map[rune]int)
	for pair, count := range pairs {
		elements[rune(pair[0])] += count
	}
	elements[rune(in[0][len(in[0])-1])] += 1
	mostCommon := []int{}
	for _, v := range elements {
		mostCommon = append(mostCommon, v)
	}
	sort.Ints(mostCommon)
	return mostCommon[len(mostCommon)-1] - mostCommon[0]
}
