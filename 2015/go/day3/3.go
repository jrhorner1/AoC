package day3

import (
	"io/ioutil"
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

var filename = "2015/input/3"

func Part1() int {
	return Puzzle(Input(filename), false)
}

func Part2() int {
	return Puzzle(Input(filename), true)
}

func Input(file string) string {
	input, _ := ioutil.ReadFile(file)
	output := strings.TrimSpace(string(input))
	return output
}

func Puzzle(input string, part2 bool) int {
	santa, robosanta := geom.Point{0, 0}, geom.Point{0, 0}
	visited := make(map[geom.Point]int)
	visited[santa], visited[robosanta] = 1, 1
	// for each direction move, then add it to the map. if already visited, the count increases
	for i, direction := range input {
		if part2 {
			if i%2 == 0 { // santa moves on odd (even index) directions
				movePoint(&santa, direction)
				visited[santa] += 1
			} else {
				movePoint(&robosanta, direction)
				visited[robosanta] += 1
			}
		} else {
			movePoint(&santa, direction)
			visited[santa] += 1
		}
	}
	return len(visited)
}

func movePoint(p *geom.Point, dir rune) {
	switch dir {
	case '^': // north
		p.Y++
	case 'v': // south
		p.Y--
	case '>': // east
		p.X++
	case '<': // west
		p.X--
	}
}
