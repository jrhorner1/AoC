package day2

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var filename = "2017/input/2"

func Part1() int {
	return Puzzle(Input(filename), false)
}

func Part2() int {
	return Puzzle(Input(filename), true)
}

func Input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func Puzzle(input []string, part2 bool) int {
	diffs, divs := []int{}, []int{}
	for _, row := range input {
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
