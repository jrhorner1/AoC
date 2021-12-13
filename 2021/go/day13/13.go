package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math/geometry"
)

type fold struct {
	axis  string
	point int
}

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n\n")
	dots := make(map[geometry.Point]interface{})
	for _, line := range strings.Split(strings.TrimSpace(in[0]), "\n") {
		l := strings.Split(line, ",")
		x, _ := strconv.Atoi(l[0])
		y, _ := strconv.Atoi(l[1])
		dots[geometry.Point{X: x, Y: y}] = struct{}{}
	}
	folds := []fold{}
	for _, line := range strings.Split(strings.TrimSpace(in[1]), "\n") {
		l := strings.Fields(line)
		f := strings.Split(l[2], "=")
		i, _ := strconv.Atoi(f[1])
		folds = append(folds, fold{f[0], i})
	}
	for _, f := range folds {
		if f.axis == "x" {
			for dot, _ := range dots {
				if dot.X > f.point {
					new := geometry.Point{dot.X - (dot.X-f.point)*2, dot.Y}
					delete(dots, dot)
					dots[new] = struct{}{}
				}
			}
		}
		if f.axis == "y" {
			for dot, _ := range dots {
				if dot.Y > f.point {
					new := geometry.Point{dot.X, dot.Y - (dot.Y-f.point)*2}
					delete(dots, dot)
					dots[new] = struct{}{}
				}
			}
		}
		if !part2 {
			break
		}
	}
	if part2 {
		printDots(dots, 39, 6)
	}
	return len(dots)
}

func printDots(dots map[geometry.Point]interface{}, xsize, ysize int) {
	for y := 0; y < ysize; y++ {
		for x := 0; x < xsize; x++ {
			if _, found := dots[geometry.Point{X: x, Y: y}]; found {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
