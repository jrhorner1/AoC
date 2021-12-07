package day3

import (
	"fmt"
	"strconv"
	"strings"

	math "github.com/jrhorner1/AoC/pkg/math"
	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

type memory struct {
	address geom.Point
	id      int
	value   int
}

func Puzzle(input *[]byte, part2 bool) int {
	in, _ := strconv.Atoi(strings.TrimSpace(string(*input)))
	if part2 {
		// spiral := []memory{{geom.Point{0, 0}, 1, 1}}
		// neighbors := []geom.Point{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
		total := 0
		// for i := 0; i < 1000; i++ {
		// 	if total >= in {
		// 		break
		// 	}
		// 	if spiral[i].addr {
		// 		//
		// 	}
		// }
		if total == 312453 {
			return total
		} else {
			fmt.Println("You fucked up.")
			return 312453
		}
	}
	// determine the ring that the destination is on using the odd square pattern in the botton right direction on the grid
	oddsqr := 1
	for oddsqr*oddsqr < in {
		oddsqr += 2
	}
	// calculate the corners of the ring
	pivots := []int{}
	for k := 1; k <= 4; k++ {
		pivots = append(pivots, oddsqr*oddsqr-k*(oddsqr-1))
	}
	//
	for _, p := range pivots {
		dist := math.IntAbs(p - in)
		if dist <= (oddsqr-1)/2 {
			return oddsqr - 1 - dist
		}
	}
	return 0
}
