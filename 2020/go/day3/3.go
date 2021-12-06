package day3

import (
	"image"
	"io/ioutil"
	"strings"
)

var filename = "2020/input/3"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Fields(string(input))
	return output
}

func puzzle(input []string, part2 bool) int {
	slopes := map[image.Point]int{{1, 1}: 0, {3, 1}: 0, {5, 1}: 0, {7, 1}: 0, {1, 2}: 0}
	product := 1
	for s := range slopes {
		for p := (image.Point{}); p.Y < len(input); p = p.Add(s) {
			if input[p.Y][p.X%len(input[0])] == '#' {
				slopes[s]++
			}
		}
		product *= slopes[s]
	}
	if part2 {
		return product
	}
	return slopes[image.Point{3, 1}]
}
