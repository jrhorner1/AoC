package day12

import (
	"fmt"
	"image"
	"math"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n")
	ship, neoship, waypoint := image.Point{0, 0}, image.Point{0, 0}, image.Point{10, -1}
	if part2 {
		return navigate(in, &neoship, &waypoint, &waypoint)
	}
	return navigate(in, &ship, &image.Point{1, 0}, &ship)
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
