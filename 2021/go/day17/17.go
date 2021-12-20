package day17

import (
	"math"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math/geometry"
)

func Puzzle(input *[]byte, part2 bool) int {
	target := getTarget(string(*input))
	velocityHeights, velocities := getVelocities(&target)
	maxHeight := -(math.MaxInt64)
	for _, height := range velocityHeights {
		if height > maxHeight {
			maxHeight = height
		}
	}
	if part2 {
		return len(velocities)
	}
	return maxHeight
}

func getTarget(input string) map[geometry.Point]interface{} {
	in := strings.Split(strings.TrimSpace(input), " ")
	targetXStrings := strings.Split(in[2][2:len(in[2])-1], "..")
	targetYStrings := strings.Split(in[3][2:], "..")
	targetX, targetY := []int{}, []int{}
	for _, s := range targetXStrings {
		number, _ := strconv.Atoi(s)
		targetX = append(targetX, number)
	}
	for _, s := range targetYStrings {
		number, _ := strconv.Atoi(s)
		targetY = append(targetY, number)
	}
	target := make(map[geometry.Point]interface{})
	for y := targetY[0]; y <= targetY[1]; y++ {
		for x := targetX[0]; x <= targetX[1]; x++ {
			target[geometry.Point{X: x, Y: y}] = struct{}{}
		}
	}
	return target
}

func getVelocities(target *map[geometry.Point]interface{}) (map[geometry.Point]int, []geometry.Point) {
	velMap := make(map[geometry.Point]int)
	velSlice := []geometry.Point{}
	for y := -150; y < 1000; y++ { // y is angle
		for x := 0; x < 500; x++ { // x is forward momentum
			probe := geometry.Point{X: 0, Y: 0}
			velocity := geometry.Point{X: x, Y: y}
			peak := 0
			for step := 0; step < 1000; step++ {
				probe.X += velocity.X
				probe.Y += velocity.Y
				if probe.Y > peak {
					peak = probe.Y
				}
				if _, found := (*target)[probe]; found {
					velMap[velocity] = peak
					velSlice = append(velSlice, velocity)
					break
				}
				// drag
				if velocity.X > 0 {
					velocity.X--
				} else if velocity.X < 0 {
					velocity.X++
				}
				velocity.Y-- // gravity
			}
		}
	}
	return velMap, velSlice
}
