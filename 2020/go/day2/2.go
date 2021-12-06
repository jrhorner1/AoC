package day2

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var filename = "2020/input/2"

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
	p1, p2 := 0, 0
	for _, i := range input {
		var min, max int
		var char byte
		var pw string
		fmt.Sscanf(i, "%v-%v %c: %v", &min, &max, &char, &pw)
		count := strings.Count(pw, string(char))
		if count >= min && count <= max {
			p1++
		}
		if (pw[min-1] == char) != (pw[max-1] == char) {
			p2++
		}
	}
	if part2 {
		return p2
	}
	return p1
}
