package day2

import (
	"io/ioutil"
	"strconv"
	"strings"

	geo "github.com/jrhorner1/AoC/pkg/math/geometry"
)

var filename = "2021/input/2"

type submarine struct {
	location geo.Point
	aim      int
}

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
	p1 := submarine{geo.Point{0, 0}, 0}
	p2 := submarine{geo.Point{0, 0}, 0}
	for _, i := range input {
		instruction := strings.Fields(i)
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1])
		if direction == "forward" {
			p1.location.X += distance          // move forward
			p2.location.X += distance          // move forward
			p2.location.Y += p2.aim * distance // adjust depth using aim * distance
		} else if direction == "up" {
			p1.location.Y -= distance // move higher
			p2.aim -= distance        // adjust aim upward
		} else if direction == "down" {
			p1.location.Y += distance // move deeper
			p2.aim += distance        // adjust aim downward
		}
	}
	if part2 {
		return p2.location.X * p2.location.Y
	}
	return p1.location.X * p1.location.Y
}
