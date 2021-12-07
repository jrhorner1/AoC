package day3

import (
	"image"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Fields(string(*input))
	slopes := map[image.Point]int{{1, 1}: 0, {3, 1}: 0, {5, 1}: 0, {7, 1}: 0, {1, 2}: 0}
	product := 1
	for s := range slopes {
		for p := (image.Point{}); p.Y < len(in); p = p.Add(s) {
			if in[p.Y][p.X%len(in[0])] == '#' {
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
