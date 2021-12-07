package day3

import (
	"io/ioutil"
	"strconv"
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

var filename = "2019/input/3"

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
	redWire := getPath(strings.Split(strings.TrimSpace(string(input[0])), ","))
	greenWire := getPath(strings.Split(strings.TrimSpace(string(input[1])), ","))
	if part2 {
		return getBestDistance(false, &redWire, &greenWire)
	}
	return getBestDistance(true, &redWire, &greenWire)
}

func getPath(wire []string) []geom.Point {
	var position = geom.Point{X: 0, Y: 0}
	var path []geom.Point
	path = append(path, position)
	for i := range wire {
		instruction := wire[i]
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])
		switch direction {
		case 'R':
			for ; distance > 0; distance-- {
				position.X++
				path = append(path, position)
			}
		case 'L':
			for ; distance > 0; distance-- {
				position.X--
				path = append(path, position)
			}
		case 'U':
			for ; distance > 0; distance-- {
				position.Y++
				path = append(path, position)
			}
		case 'D':
			for ; distance > 0; distance-- {
				position.Y--
				path = append(path, position)
			}
		}

	}
	return path
}

func getBestDistance(manhatten bool, red *[]geom.Point, green *[]geom.Point) int {
	min, last := 0, 0
	for stepA, i := range *red {
		for stepB, j := range *green {
			if i.X == j.X && i.Y == j.Y {
				if manhatten {
					last = (*red)[0].ManhattanDistance(i)
				} else {
					last = stepA + stepB
				}
				if min == 0 || last < min {
					min = last
				}
			}
		}
	}
	return min
}
