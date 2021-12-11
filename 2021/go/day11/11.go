package day11

import (
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math/geometry"
)

type dumboOctopus struct {
	position geometry.Point
	energy   int
	flashed  bool
}

func (octopus *dumboOctopus) charge() { octopus.energy++ }
func (octopus *dumboOctopus) flash()  { octopus.flashed, octopus.energy = true, 0 }
func (octopus *dumboOctopus) reset()  { octopus.flashed = false }

func Puzzle(input *[]byte, part2 bool) int {
	octopi := []dumboOctopus{}
	for x, row := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		for y, e := range strings.TrimSpace(row) {
			position := geometry.Point{x, y}
			energy, _ := strconv.Atoi(string(e))
			octopi = append(octopi, dumboOctopus{position, energy, false})
		}
	}
	// steps := len(octopi) // 100 steps since there 10*10 octopi
	steps := len(octopi)
	if part2 {
		steps *= 10
	}
	flashes := 0
	for step := 1; step <= steps; step++ {
		for i := range octopi {
			octopi[i].charge()
		}
		for i := 0; i < len(octopi)/4; i++ {
			for j := range octopi {
				if octopi[j].energy > 9 && !octopi[j].flashed {
					octopi[j].flash()
					neighbors := []geometry.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
					for _, n := range neighbors {
						neighbor := geometry.Point{octopi[j].position.X + n.X, octopi[j].position.Y + n.Y}
						for k := range octopi {
							if octopi[k].position.X == neighbor.X && octopi[k].position.Y == neighbor.Y {
								if !octopi[k].flashed {
									octopi[k].charge()
								}
							}
						}
					}
					flashes++
				}
			}
		}
		synced := true
		for i := range octopi {
			if octopi[i].flashed {
				octopi[i].reset()
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
