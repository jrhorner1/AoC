package day3

import (
	"strconv"
	"strings"

	math "github.com/jrhorner1/AoC/pkg/math"
	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

func Puzzle(input *[]byte, part2 bool) int {
	target, _ := strconv.Atoi(strings.TrimSpace(string(*input)))
	if part2 {
		current := geom.Point{0, 0}
		spiral := make(map[geom.Point]int)
		spiral[current] = 1
		neighbors := []geom.Point{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
		for {
			if spiral[current] >= target {
				return spiral[current]
			}
			if current.X-current.Y >= 0 && current.X+current.Y <= 0 { // right
				current.X++
			} else if current.X-current.Y > 0 && current.X+current.Y >= 0 { // up
				current.Y++
			} else if current.X-current.Y <= 0 && current.X+current.Y > 0 { // left
				current.X--
			} else if current.X-current.Y <= 0 && current.X+current.Y <= 0 { // down
				current.Y--
			}
			for _, n := range neighbors {
				neighbor := geom.Point{current.X + n.X, current.Y + n.Y}
				if v, found := spiral[neighbor]; found {
					spiral[current] += v
				}
			}
		}
	}
	// determine the ring that the destination is on using the odd square pattern in the botton right direction on the grid
	oddsqr := 1
	for oddsqr*oddsqr < target {
		oddsqr += 2
	}
	// calculate the corners of the ring
	pivots := []int{}
	for k := 1; k <= 4; k++ {
		pivots = append(pivots, oddsqr*oddsqr-k*(oddsqr-1))
	}
	//
	for _, p := range pivots {
		dist := math.IntAbs(p - target)
		if dist <= (oddsqr-1)/2 {
			return oddsqr - 1 - dist
		}
	}
	return 0
}
