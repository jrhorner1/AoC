package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type point struct {
	x int // east/west
	y int // north/south
}

func main() {
	input, _ := ioutil.ReadFile("2015/input/3")
	directions := strings.TrimSpace(string(input))
	// set santa's position
	santa := point{0, 0}
	// create a map of visited locations
	visited := make(map[point]int)
	// add santa's current location to map
	visited[santa] = 1
	// for each direction move, then add it to the map. if already visited, the count increases
	for _, direction := range directions {
		movePoint(&santa, direction)
		visited[santa] += 1
	}
	fmt.Println("Part 1:", len(visited))

	// reset santa's position
	santa = point{0, 0}
	robosanta := point{0, 0}
	// clear the visited map for next year
	for k, _ := range visited {
		delete(visited, k)
	}
	visited[santa], visited[robosanta] = 1, 1
	for i, direction := range directions {
		if i%2 == 0 { // santa moves on even index directions
			movePoint(&santa, direction)
			visited[santa] += 1
		} else {
			movePoint(&robosanta, direction)
			visited[robosanta] += 1
		}
	}
	fmt.Println("Part 2:", len(visited))
}

func movePoint(p *point, dir rune) {
	switch dir {
	case '^': // north
		p.y++
	case 'v': // south
		p.y--
	case '>': // east
		p.x++
	case '<': // west
		p.x--
	}
}
