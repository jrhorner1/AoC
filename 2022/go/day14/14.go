package day14

import (
	"fmt"
	"image"
	"strings"

	"github.com/sirupsen/logrus"
)

var sandStart = image.Point{X: 500, Y: 0}

func Puzzle(input *[]byte, part2 bool) int {
	//logrus.SetLevel(logrus.DebugLevel)
	cave := make(Cave)
	lowestPoint := 0
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		start := image.Point{}
		logrus.Debug(len(strings.Split(strings.TrimSpace(line), "->")))
		for i, coord := range strings.Split(strings.TrimSpace(line), "->") {
			if i == 0 {
				start = parsePoint(coord)
				cave[start] = nil
				continue
			}
			dest := parsePoint(coord)
			logrus.Debug(start, dest)
			if start.X == dest.X {
				if start.Y < dest.Y {
					cave.populate(start.Y, dest.Y, dest.X, "up")
				} else {
					cave.populate(dest.Y, start.Y, dest.X, "down")
				}
			} else if start.Y == dest.Y {
				if start.X < dest.X {
					cave.populate(start.X, dest.X, dest.Y, "right")
				} else {
					cave.populate(dest.X, start.X, dest.Y, "left")
				}
			}
			if start.Y > lowestPoint {
				lowestPoint = start.Y
			}
			if dest.Y > lowestPoint {
				lowestPoint = dest.Y
			}
			start = dest
		}
	}
	if part2 {
		cave.populate(sandStart.X-sandStart.X, sandStart.X*2, lowestPoint+2, "right")
	}
	logrus.Debug(len(cave), cave)
	totalSand := 0
	last := image.Point{}
	sand := sandStart
	for {
		for i, delta := range []image.Point{{0, 1}, {-1, 1}, {1, 1}} {
			if _, exists := cave[sand.Add(delta)]; !exists {
				last = sand
				sand = sand.Add(delta)
				break
			} else {
				if i == 2 { // all possible locations already occupied
					totalSand++
					cave[sand] = nil
					last = sand
					sand = sandStart
				}
			}
		}
		if part2 {
			if last == sand {
				break
			}
		} else if last.Y >= lowestPoint {
			break
		}
	}
	return totalSand
}

func parsePoint(coord string) (point image.Point) {
	_, err := fmt.Sscanf(coord, "%d,%d", &point.X, &point.Y)
	if err != nil {
		logrus.Error(err)
	}
	return
}

type Cave map[image.Point]any

func (cave Cave) populate(start, end, fixed int, direction string) {
	for i := start; i <= end; i++ {
		if direction == "up" || direction == "down" {
			cave[image.Point{X: fixed, Y: i}] = nil
		} else if direction == "right" || direction == "left" {
			cave[image.Point{X: i, Y: fixed}] = nil
		}
	}
}
