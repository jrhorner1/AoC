package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2020/input/3")
	grid := strings.Fields(string(input))
	slopes := map[image.Point]int{{1, 1}: 0, {3, 1}: 0, {5, 1}: 0, {7, 1}: 0, {1, 2}: 0}
	product := 1
	for s := range slopes {
		for p := (image.Point{}); p.Y < len(grid); p = p.Add(s) {
			if grid[p.Y][p.X%len(grid[0])] == '#' {
				slopes[s]++
			}
		}
		product *= slopes[s]
	}
	fmt.Println("Part 1:", slopes[image.Point{3, 1}])
	fmt.Println("Part 2:", product)
}
