package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	geo "github.com/jrhorner1/AoC/pkg/math/geometry"
)

type submarine struct {
	location geo.Point
	aim      int
}

func main() {
	input, _ := ioutil.ReadFile("2021/input/2")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	p1 := submarine{geo.Point{0, 0}, 0}
	p2 := submarine{geo.Point{0, 0}, 0}
	for _, i := range in {
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
		} else {
			fmt.Println("Unknown direction")
		}
	}

	fmt.Println("Part 1:", p1.location.X*p1.location.Y)
	fmt.Println("Part 2:", p2.location.X*p2.location.Y)
	fmt.Println("Happy Holidays 2021!")
}
