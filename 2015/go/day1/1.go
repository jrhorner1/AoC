package day1

import (
	"io/ioutil"
	"strings"
)

var filename = "2015/input/1"

func Part1() int {
	return Puzzle(Input(filename), false)
}

func Part2() int {
	return Puzzle(Input(filename), true)
}

func Input(file string) string {
	input, _ := ioutil.ReadFile(file)
	output := strings.TrimSpace(string(input))
	return output
}

func Puzzle(input string, part2 bool) int {
	floor, firstBasementVisit := 0, 0
	for i, c := range input {
		switch c {
		case '(': // up
			floor++
		case ')': // down
			floor--
		default:
			continue
		}
		if part2 && floor < 0 && firstBasementVisit == 0 {
			// slice starts at 0 and instructions state first position is 1
			return i + 1
		}
	}
	return floor
}
