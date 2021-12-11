package day11

import (
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math/geometry"
)

type dumboOctopus struct {
	energy  int
	flashed bool
}

func (octopus *dumboOctopus) charge() dumboOctopus { octopus.energy++; return *octopus }
func (octopus *dumboOctopus) flash() dumboOctopus {
	octopus.flashed, octopus.energy = true, 0
	return *octopus
}
func (octopus *dumboOctopus) reset() dumboOctopus { octopus.flashed = false; return *octopus }

func Puzzle(input *[]byte, steps int) int {
	octopi := map[geometry.Point]dumboOctopus{}
	for x, row := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		for y, e := range strings.TrimSpace(row) {
			position := geometry.Point{x, y}
			energy, _ := strconv.Atoi(string(e))
			octopi[position] = dumboOctopus{energy, false}
		}
	}
	flashes := 0
	for step := 1; step <= steps; step++ {
		for position, octopus := range octopi {
			octopi[position] = octopus.charge()
		}
		for i := 0; i < len(octopi)/4; i++ {
			for position, octopus := range octopi {
				if octopus.energy > 9 && !octopus.flashed {
					octopi[position] = octopus.flash()
					neighbors := []geometry.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
					for _, n := range neighbors {
						neighborPosition := geometry.Point{position.X + n.X, position.Y + n.Y}
						if neighbor, found := octopi[neighborPosition]; found {
							if !neighbor.flashed {
								octopi[neighborPosition] = neighbor.charge()
							}
						}
					}
					flashes++
				}
			}
		}
		synced := true
		for position, octopus := range octopi {
			if octopus.flashed {
				octopi[position] = octopus.reset()
			} else {
				synced = false
			}
		}
		if synced {
			return step
		}
	}
	return flashes
}
