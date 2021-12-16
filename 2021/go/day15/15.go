package day15

import (
	"math"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math/geometry"
)

func Puzzle(input *[]byte, part2 bool) int {
	cave := make(map[geometry.Point]int)
	Y, X := 0, 0
	for y, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		for x, r := range line {
			risk, _ := strconv.Atoi(string(r))
			cave[geometry.Point{X: x, Y: y}] = risk
			X = x + 1
		}
		Y = y + 1
	}
	if part2 {
		tileMultiplier := 5
		for y := 0; y < Y*tileMultiplier; y++ {
			for x := 0; x < X*tileMultiplier; x++ {
				cave[geometry.Point{x, y}] = (((cave[geometry.Point{x % X, y % Y}] + (x / X) + (y / Y)) - 1) % 9) + 1
			}
		}
		risk := dijkstra(cave, geometry.Point{0, 0})
		return risk[geometry.Point{(X * tileMultiplier) - 1, (Y * tileMultiplier) - 1}]
	}
	risk := dijkstra(cave, geometry.Point{0, 0})
	return risk[geometry.Point{X - 1, Y - 1}]
}

type queue []geometry.Point

func (q *queue) dequeue(p geometry.Point) {
	for i, qp := range *q {
		if p == qp {
			(*q)[i] = (*q)[len(*q)-1]
			*q = (*q)[:len(*q)-1]
		}
	}
}

func dijkstra(graph map[geometry.Point]int, source geometry.Point) map[geometry.Point]int {
	risk := make(map[geometry.Point]int)
	queue := queue{}
	risk[source] = 0
	for point, _ := range graph {
		if point != source {
			risk[point] = math.MaxInt64 / 2
		}
		queue = append(queue, point)
	}
	for len(queue) != 0 {
		point, min := geometry.Point{}, math.MaxInt64
		for _, p := range queue {
			if risk[p] < min {
				min = risk[p]
				point = p
			}
		}
		queue.dequeue(point)
		neighbors := []geometry.Point{{point.X - 1, point.Y}, {point.X, point.Y - 1}, {point.X + 1, point.Y}, {point.X, point.Y + 1}}
		for _, neighbor := range neighbors {
			for _, inQueue := range queue {
				if neighbor == inQueue {
					alt := risk[point] + graph[neighbor]
					if alt < risk[neighbor] {
						risk[neighbor] = alt
					}
				}
			}
		}
	}
	return risk
}
