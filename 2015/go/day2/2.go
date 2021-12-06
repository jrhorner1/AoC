package day2

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var filename = "2015/input/2"

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
	wrappingPaper, ribbon := 0, 0
	for _, d := range input {
		dimensions := strings.Split(d, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		a := 2*l*w + 2*w*h + 2*h*l
		s := []int{l * w, w * h, h * l}
		sort.Ints(s)
		p := []int{2*l + 2*w, 2*w + 2*h, 2*h + 2*l}
		sort.Ints(p)
		v := l * w * h

		wrappingPaper += a + s[0]
		ribbon += p[0] + v
	}
	if part2 {
		return ribbon
	}
	return wrappingPaper
}
