package day1

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math"
	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

var filename = "2017/input/1"

func Part1() int {
	return Puzzle(Input(filename), false)
}

func Part2() int {
	return Puzzle(Input(filename), true)
}

func Input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), ", ")
	return output
}

func Puzzle(input []string, part2 bool) int {
	path := []geom.Point{{0, 0}} // origin point
	facing := 'N'                // start off facing North
	position := geom.Point{0, 0}
	for _, d := range input {
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
	if part2 {
		firstDoubleVisit := geom.Point{0, 0}
	loop:
		for i := 1; i < len(path); i++ {
			for j := 0; j < i; j++ {
				if (path[i].X == path[j].X) && (path[i].Y == path[j].Y) {
					firstDoubleVisit = path[i]
					break loop
				}
			}
		}
		return math.IntAbs(firstDoubleVisit.X) + math.IntAbs(firstDoubleVisit.X)
	}
	return math.IntAbs(destination.X) + math.IntAbs(destination.Y)

}

func north(facing *rune, distance int, position *geom.Point, path *[]geom.Point) {
	*facing = 'N'
	for ; distance > 0; distance-- {
		(*position).X++
		*path = append(*path, *position)
	}
}
func east(facing *rune, distance int, position *geom.Point, path *[]geom.Point) {
	*facing = 'E'
	for ; distance > 0; distance-- {
		(*position).Y++
		*path = append(*path, *position)
	}
}
func west(facing *rune, distance int, position *geom.Point, path *[]geom.Point) {
	*facing = 'W'
	for ; distance > 0; distance-- {
		(*position).Y--
		*path = append(*path, *position)
	}
}
func south(facing *rune, distance int, position *geom.Point, path *[]geom.Point) {
	*facing = 'S'
	for ; distance > 0; distance-- {
		(*position).X--
		*path = append(*path, *position)
	}
}
