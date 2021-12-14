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
		pairs[pair] += 1 // count number of instances of each pair. this also gives an accurate count of either the first or second element
	}
	rules := make(map[string]byte)
	for _, line := range strings.Split(strings.TrimSpace(in[1]), "\n") {
		l := strings.Split(line, " -> ")
		rules[l[0]] = l[1][0] // accessing elements in a string like a slice results in a byte instead of rune
	}
	for round := rounds; round > 0; round-- { // number of rounds is passed into the function
		newPairs := make(map[string]int) // insertion happens all at once. make all changes in a new map, then write at the end of the round
		for pair, count := range pairs { // loop through the pairs
			if insert, found := rules[pair]; found { // check for pair insertion rule
				newPairs[string([]byte{pair[0], insert})] += count
				newPairs[string([]byte{insert, pair[1]})] += count
			} else { // if theres no rule
				newPairs[pair] = count // add the pair to the new map with the current count
			}
		}
		pairs = newPairs // write the new pairs map
	}
	elements := make(map[rune]int)
	for pair, count := range pairs { // use count for first element in each pair
		elements[rune(pair[0])] += count // since each pair overlaps this excludes the last element from the template
	}
	elements[rune(in[0][len(in[0])-1])] += 1 // increment for the last element from the template
	mostCommon := []int{}
	for _, v := range elements {
		mostCommon = append(mostCommon, v) // put the value of each element in a slice for sorting
	}
	sort.Ints(mostCommon) // sort the slice to get most/least common elements
	return mostCommon[len(mostCommon)-1] - mostCommon[0]
}
