package day12

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strings"
)

var filename = "2020/input/12"

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
	ship, neoship, waypoint := image.Point{0, 0}, image.Point{0, 0}, image.Point{10, -1}
	if part2 {
		return navigate(input, &neoship, &waypoint, &waypoint)
	}
	return navigate(input, &ship, &image.Point{1, 0}, &ship)
}

func navigate(in []string, ship, facing, move *image.Point) int {
	action := map[rune]image.Point{'N': {0, -1}, 'E': {1, 0}, 'W': {-1, 0}, 'S': {0, 1}, 'L': {-1, 1}, 'R': {1, -1}}
	for i, s := range in {
		var dir rune
		var val int
		fmt.Sscanf(s, "%c%d", &dir, &val)
		switch dir {
		case 'N', 'E', 'W', 'S':
			*move = move.Add(action[dir].Mul(val))
		case 'L', 'R':
			for i = 0; i < val/90; i++ {
				facing.X, facing.Y = action[dir].Y*facing.Y, action[dir].X*facing.X
			}
		case 'F':
			*ship = ship.Add(facing.Mul(val))
		}
	}
	return int(math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y)))
}
