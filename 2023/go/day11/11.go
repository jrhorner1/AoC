package day11

import (
	"strings"

	geo "github.com/jrhorner1/AoC/pkg/math/geometry"
	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, scale int) int {
	scale--
	log.Debug("Scale: ", scale)
	universe, maxX, maxY := newUniverse(input)
	universe.expand(maxY, false, &scale)
	universe.expand(maxX, true, &scale)
	// calculate manhattan distance between each galaxy
	sum := 0
	for g1 := range *universe {
		for g2 := range *universe {
			sum += g1.ManhattanDistance(g2)
		}
		delete(*universe, g1)
	}
	return sum
}

type Galaxy geo.Point

func newGalaxy(x, y int) *Galaxy {
	return &Galaxy{X: x, Y: y}
}

func (g *Galaxy) ManhattanDistance(q Galaxy) int {
	g1p, g2p := geo.Point(*g), geo.Point(q)
	return g1p.ManhattanDistance(g2p)
}

type Universe map[Galaxy]any

func newUniverse(input *[]byte) (u *Universe, maxX *int, maxY *int) {
	l := strings.Split(strings.TrimSpace(string(*input)), "\n")
	*maxX, *maxY = len((l)[0]), len(l)
	for y, line := range l {
		for x, r := range line {
			if r == '#' {
				(*u)[*newGalaxy(x, y)] = nil
			}
		}
	}
	log.Debug(u)
	return u, maxX, maxY
}

func (u *Universe) expand(maxP *int, isX bool, scale *int) {
expand:
	for p := 0; p < *maxP; p++ {
		// do any galaxies exist on this plane?
		for g := range *u {
			c := g.Y
			if isX {
				c = g.X
			}
			if c == p {
				continue expand
			}
		}
		// if not, move every galaxy after this plane by scale
		nu := Universe{}
		for g := range *u {
			if isX {
				if g.X > p {
					g.X += *scale
				}
			} else {
				if g.Y > p {
					g.Y += *scale
				}
			}
			nu[g] = nil
		}
		*u = nu
		*maxP += *scale // max space expanded
		p += *scale     // bypass newly expanded space
	}
	log.Debug(*u)
}
