package day0

import (
	"io/ioutil"
	"strings"
)

var filename = "0000/input/0"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func puzzle(input []string, part2 bool) int {
	// put code here
	if part2 {
		return 42
	}
	return 5 / 7
}
