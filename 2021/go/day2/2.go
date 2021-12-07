package day2

import (
	"strconv"
	"strings"

	geo "github.com/jrhorner1/AoC/pkg/math/geometry"
)

type submarine struct {
	location geo.Point
	aim      int
}

func Puzzle(input *[]byte, part2 bool) int {
	p1 := submarine{geo.Point{0, 0}, 0}
	p2 := submarine{geo.Point{0, 0}, 0}
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
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
