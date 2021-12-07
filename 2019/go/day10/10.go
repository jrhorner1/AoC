package day10

import (
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

var filename = "2019/input/10"

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

type Asteroid struct {
	x, y, m, au float64
	los         int
	zapped      bool
}

func puzzle(input []string, part2 bool) int {
	grid := GetGrid(input)
	station := Asteroid{x: 0, y: 0, los: 0}
	for _, asteroid := range *grid {
		if asteroid.los > station.los {
			station = asteroid
		}
	}
	if part2 {
		var slopes []float64
	loop:
		for i := range *grid { // get the angle of each asteroid from the station
			if station != (*grid)[i] {
				(*grid)[i].au = GetAU(station, (*grid)[i])
				(*grid)[i].m = GetSlope(station, (*grid)[i])
				for j := range slopes {
					if slopes[j] == (*grid)[i].m {
						continue loop
					}
				}
				slopes = append(slopes, (*grid)[i].m)
			}
		}
		sort.Float64s(slopes) // sort slopes numerically so we can simulate the rotating laser
		// fmt.Println(slopes)
		var start int
		for index, slope := range slopes {
			if slope == 0 { // set the starting point for the laser to slope 0 (up in the grid input)
				start = index
			}
		}
		var zapOrder []Asteroid
		for i := start; i >= 0; i-- { // to rotate clockwise on the grid, we need to count down for the right side of the grid
			asteroidsLeft := 0
			for _, asteroid := range *grid { // loop through the grid to determine how many asteroids are left
				if !asteroid.zapped {
					asteroidsLeft++
				}
			}
			if asteroidsLeft == 1 { // break the loop where there is only 1 asteroid
				break // left: the asteroid that the station is on
			}
			candidates := make([]Asteroid, 0)
			for _, asteroid := range *grid { // generate a list of candidates for vaporization
				if asteroid != station && asteroid.m == slopes[i] && !asteroid.zapped {
					candidates = append(candidates, asteroid)
				}
			}
			sort.Slice(candidates, func(i, j int) bool { // sort the candidates by astronomical units (distance) from the station
				return candidates[i].au < candidates[j].au
			})
			if len(candidates) > 0 { // if the list of candidates is not zero,
				for k, asteroid := range *grid { // loop through the grid to
					if asteroid == candidates[0] { // get its index and
						(*grid)[k].zapped = true // vaporize it
						zapOrder = append(zapOrder, candidates[0])
					}
				}
			}
			if i == 0 { // reset the counter to continue the rotation from down to left on the grid
				i = len(slopes) // set this to len(slopes) because the loop will subtract 1
			}
		}
		return int(zapOrder[199].x*100 + zapOrder[199].y)
	}
	return station.los
}

func GetSlope(a1, a2 Asteroid) float64 {
	m := math.Atan2(a1.x-a2.x, a1.y-a2.y)
	return m
}

func GetGrid(input []string) *[]Asteroid {
	var grid []Asteroid
	y := 0
	for _, in := range input {
		line := strings.Split(in, "")
		for x := range line {
			if line[x] == "#" {
				asteroidCoord := Asteroid{x: float64(x), y: float64(y), los: 0, m: 0, au: 1}
				grid = append(grid, asteroidCoord)
			}
		}
		y++
	}
	for i := range grid { // calculate asteroid line of sight
		var slopes []float64
		count := 0
	loop:
		for j := range grid {
			if grid[i] != grid[j] {
				m := GetSlope(grid[i], grid[j])
				for _, slope := range slopes {
					if slope == m {
						continue loop
					}
				}
				slopes = append(slopes, m)
				count++
			}
		}
		grid[i].los = count // set line of sight value for each asteroid
	}
	return &grid
}

func GetAU(a1, a2 Asteroid) float64 {
	au := math.Abs(a1.x-a2.x) + math.Abs(a1.y-a2.y)
	return au
}
