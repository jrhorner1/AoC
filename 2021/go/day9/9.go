package day9

import (
	"sort"
	"strconv"
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n")
	cave := make(map[geom.Point]int)
	xLen, yLen := 0, len(in)
	for y, line := range in {
		xLen = len(line)
		for x, char := range line {
			z, _ := strconv.Atoi(string(char))
			cave[geom.Point{x, y}] = z
		}
	}
	neighbors := []geom.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	lowPoints := []geom.Point{}
	totalRisk := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			low := true
			for _, neighbor := range neighbors {
				if nHeight, found := cave[geom.Point{neighbor.X + x, neighbor.Y + y}]; found {
					if nHeight <= cave[geom.Point{x, y}] {
						low = false
					}
				}
			}
			if low {
				lowPoints = append(lowPoints, geom.Point{x, y})
				totalRisk += cave[geom.Point{x, y}] + 1
			}
		}
	}
	if part2 {
		basins := []int{}
		for _, lowPoint := range lowPoints {
			size := basinSize(lowPoint, &cave, xLen, yLen)
			basins = append(basins, size)
		}
		sort.Ints(basins)
		return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
	}
	return totalRisk
}

// https://en.wikipedia.org/wiki/Flood_fill#Stack-based_recursive_implementation_(four-way)
func basinSize(start geom.Point, cave *map[geom.Point]int, xLen, yLen int) int {
	basin := make(map[geom.Point]interface{})
	queue := []geom.Point{start}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		if (*cave)[point] < 9 {
			if _, found := basin[point]; !found {
				if point.X > 0 {
					queue = append(queue, geom.Point{point.X - 1, point.Y})
				}
				if point.X < xLen-1 {
					queue = append(queue, geom.Point{point.X + 1, point.Y})
				}
				if point.Y > 0 {
					queue = append(queue, geom.Point{point.X, point.Y - 1})
				}
				if point.Y < yLen-1 {
					queue = append(queue, geom.Point{point.X, point.Y + 1})
				}
			}
			basin[point] = struct{}{}
		}
	}
	return len(basin)
}
