package day2

import (
	"sort"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	diffs, divs := []int{}, []int{}
	for _, row := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		fields := strings.Fields(row)
		intFields := []int{}
		for _, field := range fields {
			data, _ := strconv.Atoi(field)
			intFields = append(intFields, data)
		}
		sort.Ints(intFields)
		diffs = append(diffs, intFields[len(intFields)-1]-intFields[0])
		for i, a := range intFields {
			for j, b := range intFields {
				if i == j {
					continue
				}
				if a%b == 0 {
					divs = append(divs, a/b)
				}
			}
		}
	}
	checksum := 0
	if part2 {
		for _, div := range divs {
			checksum += div
		}
		return checksum
	}
	for _, diff := range diffs {
		checksum += diff
	}
	return checksum
}
