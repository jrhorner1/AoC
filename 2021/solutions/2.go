package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x int // horizontal forward
	z int // vertical depth
	a int // aim
}

func main() {
	input, _ := ioutil.ReadFile("2021/input/2")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	p1 := point{0, 0, 0}
	p2 := point{0, 0, 0}
	for _, i := range in {
		instruction := strings.Fields(i)
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1])
		if direction == "forward" {
			p1.x += distance        // move forward
			p2.x += distance        // move forward
			p2.z += p2.a * distance // adjust depth using aim * distance
		} else if direction == "up" {
			p1.z -= distance // move higher
			p2.a -= distance // adjust aim upward
		} else if direction == "down" {
			p1.z += distance // move deeper
			p2.a += distance // adjust aim downward
		} else {
			fmt.Println("Unknown direction")
		}
	}

	fmt.Println("Part 1:", p1.x*p1.z)
	fmt.Println("Part 2:", p2.x*p2.z)
	fmt.Println("Happy Holidays 2021!")
}
