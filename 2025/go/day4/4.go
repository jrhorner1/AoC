package day4

import (
	"strings"

	geo "github.com/jrhorner1/AoC/pkg/math/geometry"
	log "github.com/sirupsen/logrus"
)

const DEBUG = true

func Puzzle(input *[]byte, part2 bool) int {
	if DEBUG {
		log.SetLevel(log.DebugLevel)
	}
	grid := map[geo.Point]bool{}
	for y, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		for x, rune := range line {
			location := geo.Point{X: x, Y: y}
			if rune == '@' {
				grid[location] = true
			}
		}
	}
	accessible := removeRolls(&grid)
	if part2 {
		removed := accessible
		for accessible != 0 {
			accessible = removeRolls(&grid)
			removed += accessible
		}
		return removed
	}
	return accessible
}

func removeRolls(grid *map[geo.Point]bool) int {
	accessible := []geo.Point{}
	deltas := geo.InitDeltas(1)
	for location, occupied := range *grid {
		if occupied {
			adjacent := 0
			for _, delta := range deltas {
				if (*grid)[location.Add(delta)] {
					adjacent++
				}
			}
			if adjacent < 4 {
				accessible = append(accessible, location)
			}
		}
	}
	for _, location := range accessible {
		(*grid)[location] = false
	}
	return len(accessible)
}
