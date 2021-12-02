package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x     int
	y     int
	steps int
}

func getPath(wire []string) []point {
	var position = point{x: 0, y: 0, steps: 0}
	var path []point
	path = append(path, position)
	for i := range wire {
		instruction := wire[i]
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])
		switch direction {
		case 82: // ascii "R"
			for j := 0; j < distance; j++ {
				position.x++
				position.steps++
				path = append(path, position)
			}
		case 76: // ascii "L"
			for j := 0; j < distance; j++ {
				position.x--
				position.steps++
				path = append(path, position)
			}
		case 85: // ascii "U"
			for j := 0; j < distance; j++ {
				position.y++
				position.steps++
				path = append(path, position)
			}
		case 68: // ascii "D"
			for j := 0; j < distance; j++ {
				position.y--
				position.steps++
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

func getBestDistance(manhatten bool, wire *[][]point) int {
	min, last := 0, 0
	start := (*wire)[0][0]
	for stepA, i := range (*wire)[0] {
		for stepB, j := range (*wire)[1] {
			if i.x == j.x && i.y == j.y {
				if manhatten {
					last = start.ManhattanDistance(i)
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
	wire := [][]point{}
	for i := range in {
		wire = append(wire, getPath(strings.Split(strings.TrimSpace(string(in[i])), ",")))
	}
	fmt.Println("Part 1:", getBestDistance(true, &wire))
	fmt.Println("Part 2:", getBestDistance(false, &wire))
}
