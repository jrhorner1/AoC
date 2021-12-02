package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int // north/south plane
	y int // east/west plane
}

func intAbs(i int) int {
	return int(math.Abs(float64(i)))
}

func north(facing *rune, distance int, position *point, path *[]point) {
	*facing = 'N'
	for ; distance > 0; distance-- {
		(*position).x++
		*path = append(*path, *position)
	}
}
func east(facing *rune, distance int, position *point, path *[]point) {
	*facing = 'E'
	for ; distance > 0; distance-- {
		(*position).y++
		*path = append(*path, *position)
	}
}
func west(facing *rune, distance int, position *point, path *[]point) {
	*facing = 'W'
	for ; distance > 0; distance-- {
		(*position).y--
		*path = append(*path, *position)
	}
}
func south(facing *rune, distance int, position *point, path *[]point) {
	*facing = 'S'
	for ; distance > 0; distance-- {
		(*position).x--
		*path = append(*path, *position)
	}
}

func main() {
	input, _ := ioutil.ReadFile("2016/input/1")
	in := strings.Split(strings.TrimSpace(string(input)), ", ") // space after comma
	path := []point{{0, 0}}                                     // origin point
	facing := 'N'                                               // start off facing North
	position := point{0, 0}
	for _, d := range in {
		direction := d[0]
		distance, _ := strconv.Atoi(d[1:])
		switch direction {
		case 'L':
			switch facing {
			case 'N':
				west(&facing, distance, &position, &path)
			case 'E':
				north(&facing, distance, &position, &path)
			case 'W':
				south(&facing, distance, &position, &path)
			case 'S':
				east(&facing, distance, &position, &path)
			}
		case 'R':
			switch facing {
			case 'N':
				east(&facing, distance, &position, &path)
			case 'E':
				south(&facing, distance, &position, &path)
			case 'W':
				north(&facing, distance, &position, &path)
			case 'S':
				west(&facing, distance, &position, &path)
			}
		}
	}
	destination := path[len(path)-1]
	fmt.Println("Part 1:", intAbs(destination.x)+intAbs(destination.y))

	firstDoubleVisit := point{0, 0}
loop:
	for i := 1; i < len(path); i++ {
		for j := 0; j < i; j++ {
			if (path[i].x == path[j].x) && (path[i].y == path[j].y) {
				firstDoubleVisit = path[i]
				break loop
			}
		}
	}
	fmt.Println("Part 2:", intAbs(firstDoubleVisit.x)+intAbs(firstDoubleVisit.y))
}
