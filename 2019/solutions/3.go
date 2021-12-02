package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func getPath(wire []string) []point {
	var position = point{x: 0, y: 0}
	var path []point
	path = append(path, position)
	for i := range wire {
		instruction := wire[i]
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])
		switch direction {
		case 'R':
			for ; distance > 0; distance-- {
				position.x++
				path = append(path, position)
			}
		case 'L':
			for ; distance > 0; distance-- {
				position.x--
				path = append(path, position)
			}
		case 'U':
			for ; distance > 0; distance-- {
				position.y++
				path = append(path, position)
			}
		case 'D':
			for ; distance > 0; distance-- {
				position.y--
				path = append(path, position)
			}
		}

	}
	return path
}

func intAbs(i int) int {
	return int(math.Abs(float64(i)))
}

func (p *point) ManhattanDistance(q point) int {
	return intAbs(q.x-p.x) + intAbs(q.y-p.y)
}

func getBestDistance(manhatten bool, red *[]point, green *[]point) int {
	min, last := 0, 0
	for stepA, i := range *red {
		for stepB, j := range *green {
			if i.x == j.x && i.y == j.y {
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

func main() {
	input, _ := ioutil.ReadFile("2019/input/3")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	redWire := getPath(strings.Split(strings.TrimSpace(string(in[0])), ","))
	greenWire := getPath(strings.Split(strings.TrimSpace(string(in[1])), ","))

	fmt.Println("Part 1:", getBestDistance(true, &redWire, &greenWire))
	fmt.Println("Part 2:", getBestDistance(false, &redWire, &greenWire))
}
