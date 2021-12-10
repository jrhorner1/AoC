package day3

import (
	"strconv"
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

const fabricSize int = 1000

type Claim struct {
	id, length, width int
	start             geom.Point
}

func Puzzle(input *[]byte, part2 bool) int {
	claims := []Claim{}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		claimFields := strings.Fields(line)
		id, _ := strconv.Atoi(claimFields[0][1:])
		startString := claimFields[2][:len(claimFields[2])-1]
		start := strings.Split(startString, ",")
		x, _ := strconv.Atoi(start[0])
		y, _ := strconv.Atoi(start[1])
		s := geom.Point{x, y}
		dimensions := strings.Split(claimFields[3], "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		claims = append(claims, Claim{id, l, w, s})
	}
	fabric := [fabricSize][fabricSize]int{}
	for _, claim := range claims {
		for x := claim.start.X; x < claim.start.X+claim.length; x++ {
			for y := claim.start.Y; y < claim.start.Y+claim.width; y++ {
				fabric[x][y]++
			}
		}
	}
	if part2 {
		for _, claim := range claims {
			valid := true
			for x := claim.start.X; x < claim.start.X+claim.length; x++ {
				for y := claim.start.Y; y < claim.start.Y+claim.width; y++ {
					if fabric[x][y] > 1 {
						valid = false
					}
				}
			}
			if valid {
				return claim.id
			}
		}
	}
	overClaimed := 0
	for x := range fabric {
		for y := range fabric {
			if fabric[x][y] > 1 {
				overClaimed++
			}
		}
	}
	return overClaimed
}
