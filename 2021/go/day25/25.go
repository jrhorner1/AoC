package day25

import (
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	floor := [][]rune{}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		floor = append(floor, []rune(line))
	}
	if part2 {
		return 42
	}
	height, width := len(floor), len(floor[0])
	step := 1
	for ; step < 1000; step++ {
		newFloor := make([][]rune, height)
		for i := range newFloor {
			newFloor[i] = make([]rune, width)
			copy(newFloor[i], floor[i])
		}
		movement := false
		for y := range floor {
			for x := range floor[y] {
				nx := (x + 1) % width
				if floor[y][x] == '>' && floor[y][nx] == '.' {
					newFloor[y][nx] = '>'
					newFloor[y][x] = '.'
					movement = true
				}
			}
		}
		for i := range floor {
			copy(floor[i], newFloor[i])
		}
		for y := range floor {
			for x := range floor[y] {
				ny := (y + 1) % height
				if floor[y][x] == 'v' && floor[ny][x] == '.' {
					newFloor[ny][x] = 'v'
					newFloor[y][x] = '.'
					movement = true
				}
			}
		}
		for i := range floor {
			copy(floor[i], newFloor[i])
		}
		if !movement {
			break
		}
	}
	return step
}
