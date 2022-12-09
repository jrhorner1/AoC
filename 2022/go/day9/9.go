package day9

import (
	"fmt"
	"image"
	"strings"

	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, knots int) int {
	rope := newRope(knots)
	rope.visited[*rope.tail] = struct{}{} // set the starting point as visited
	for _, movement := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		var direction string
		var distance int
		_, err := fmt.Sscanf(movement, "%s %d", &direction, &distance)
		if err != nil {
			logrus.Error(err)
		}
		rope.move(direction, distance)
	}
	return len(rope.visited)
}

type Rope struct {
	knots      []*image.Point
	head, tail *image.Point
	visited    map[image.Point]interface{}
}

func newRope(knots int) *Rope {
	rope := &Rope{
		knots:   []*image.Point{},
		visited: make(map[image.Point]interface{}),
	}
	for i := 0; i < knots; i++ {
		rope.knots = append(rope.knots, &image.Point{0, 0})
	}
	rope.head = rope.knots[0]
	rope.tail = rope.knots[knots-1]
	return rope
}

func (r *Rope) move(direction string, distance int) {
	delta := map[string]image.Point{
		"U":  {0, 1},   // up
		"D":  {0, -1},  // down
		"R":  {1, 0},   // right
		"L":  {-1, 0},  //left
		"RU": {1, 1},   // up right
		"LU": {-1, 1},  // up left
		"RD": {1, -1},  // down right
		"LD": {-1, -1}, // down left
	}
	for travelled := 0; travelled < distance; travelled++ {
		for i, knot := range r.knots {
			if knot == r.head {
				*knot = knot.Add(delta[direction])
			} else {
				diff := r.knots[i-1].Sub(*knot)
				switch diff {
				case delta["U"].Mul(2): // up
					*knot = knot.Add(delta["U"])
				case delta["D"].Mul(2): // down
					*knot = knot.Add(delta["D"])
				case delta["L"].Mul(2): // left
					*knot = knot.Add(delta["L"])
				case delta["R"].Mul(2): // right
					*knot = knot.Add(delta["R"])
				case delta["RU"].Add(delta["U"]), delta["RU"].Add(delta["R"]), delta["RU"].Mul(2): // up right
					*knot = knot.Add(delta["RU"])
				case delta["LU"].Add(delta["U"]), delta["LU"].Add(delta["L"]), delta["LU"].Mul(2): // up left
					*knot = knot.Add(delta["LU"])
				case delta["RD"].Add(delta["D"]), delta["RD"].Add(delta["R"]), delta["RD"].Mul(2): // down right
					*knot = knot.Add(delta["RD"])
				case delta["LD"].Add(delta["D"]), delta["LD"].Add(delta["L"]), delta["LD"].Mul(2): // down left
					*knot = knot.Add(delta["LD"])
				}
			}
			if knot == r.tail {
				r.visited[*knot] = struct{}{}
			}
		}
	}
}
