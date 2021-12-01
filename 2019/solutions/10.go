package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
	"sort"
)

type Asteroid struct {
	x,y,m,au float64
	los int
	zapped bool
}

func GetSlope(a1, a2 Asteroid, ) float64 {
	m := math.Atan2(a1.x - a2.x, a1.y - a2.y)
	return m
}

func GetGrid(filename string) *[]Asteroid {
    file, _ := os.Open(filename)
    defer file.Close()
    var grid []Asteroid
    scanner := bufio.NewScanner(file)
    y := 0
    for scanner.Scan() {
        line := strings.Split(scanner.Text(),"")
        for x := range line {
        	if line[x] == "#" {
		        asteroidCoord := Asteroid{ x: float64(x), y: float64(y), los: 0, m: 0, au: 1}
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
	// fmt.Println(grid)
    return &grid
}

func GetAU(a1, a2 Asteroid, ) float64 {
	au := math.Abs(a1.x - a2.x) + math.Abs(a1.y - a2.y)
	return au
}

func Silver() (*[]Asteroid, Asteroid) {
	grid := GetGrid("input")
	// grid := GetGrid("example")
	station := Asteroid{x: 0, y: 0, los: 0}
	for _, asteroid := range *grid {
		// comment for example
		if asteroid.los > station.los {
			station = asteroid
		}
		// uncomment for example
		// if asteroid.x == 8 && asteroid.y == 3 {
		// 	station = asteroid
		// }
	}
	return grid, station
}

func Gold(grid *[]Asteroid, station Asteroid) Asteroid {
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
			break 				// left: the asteroid that the station is on
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
		if len(candidates) > 0 { 	// if the list of candidates is not zero, 
			for k, asteroid := range *grid { 	// loop through the grid to
				if asteroid == candidates[0] { 	// get its index and
					(*grid)[k].zapped = true 	// vaporize it
					zapOrder = append(zapOrder, candidates[0])
				}
			}
		}
		if i == 0 { // reset the counter to continue the rotation from down to left on the grid
			i = len(slopes) // set this to len(slopes) because the loop will subtract 1
		}
	}
	return zapOrder[199]
}

func main() {
	grid, station := Silver()
	fmt.Println("Station location:",station.x,station.y)
	fmt.Println("Asteroids in line of sight:",station.los)
	winningBet := Gold(grid, station)
	fmt.Println("Winning bet:",winningBet.x * 100 + winningBet.y)
}