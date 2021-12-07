package day3

import (
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

func Puzzle(input *[]byte, part2 bool) int {
	inputString := strings.TrimSpace(string(*input))
	santa, robosanta := geom.Point{0, 0}, geom.Point{0, 0}
	visited := make(map[geom.Point]int)
	visited[santa], visited[robosanta] = 1, 1
	// for each direction move, then add it to the map. if already visited, the count increases
	for i, direction := range inputString {
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
