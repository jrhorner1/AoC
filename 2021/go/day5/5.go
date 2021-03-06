package day5

import (
	"strconv"
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

func Puzzle(input *[]byte, part2 bool) int {
	dangerZoneHV := make(map[geom.Point]int)
	dangerZoneHVD := make(map[geom.Point]int)
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		points := strings.Split(strings.TrimSpace(i), "->")
		a := parsePoint(strings.Split(strings.TrimSpace(points[0]), ","))
		b := parsePoint(strings.Split(strings.TrimSpace(points[1]), ","))
		vent := geom.Line{a, b}
		if vent.IsHorizontal() {
			p := vent.Start()
			for i := 0; i <= vent.Length(); i++ {
				dangerZoneHV[p] += 1
				dangerZoneHVD[p] += 1
				p.X++
			}
		} else if vent.IsVertical() {
			p := vent.Start()
			for i := 0; i <= vent.Length(); i++ {
				dangerZoneHV[p] += 1
				dangerZoneHVD[p] += 1
				p.Y++
			}
		} else if vent.IsDiagonal() {
			p := vent.Start()
			for i := 0; i <= vent.Length(); i++ {
				dangerZoneHVD[p] += 1
				p.X++
				if vent.End().Y > p.Y {
					p.Y++
				} else {
					p.Y--
				}
			}
		}
	}
	if part2 {
		return getMostDangerous(&dangerZoneHVD)
	}
	return getMostDangerous((&dangerZoneHV))
}

func parsePoint(p []string) geom.Point {
	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])
	return geom.Point{x, y}
}

func getMostDangerous(dangerZone *map[geom.Point]int) int {
	mostDangerousPoints := 0
	for _, v := range *dangerZone {
		if v > 1 {
			mostDangerousPoints++
		}
	}
	return mostDangerousPoints
}
