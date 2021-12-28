package day22

import (
	"fmt"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math"
)

type cuboid struct {
	xMin, xMax, yMin, yMax, zMin, zMax int
	action                             string
	volume                             int
}

type cuboids []cuboid

func Puzzle(input *[]byte, initializing bool) int {
	var cubes cuboids
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		var cube cuboid
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &cube.action, &cube.xMin, &cube.xMax, &cube.yMin, &cube.yMax, &cube.zMin, &cube.zMax)
		cube.volume = cube.getVolume()
		cubes = append(cubes, cube)
	}
	if initializing {
		init := cuboids{}
		initMin, initMax := -50, 50
		for _, cube := range cubes {
			if cube.xMin >= initMin && cube.xMax <= initMax && cube.yMin >= initMin && cube.yMax <= initMax && cube.zMin >= initMin && cube.zMax <= initMax {
				init = append(init, cube)
			}
		}
		return init.count()
	}
	return cubes.count()
}

func (c *cuboid) getVolume() int {
	return (c.xMax + 1 - c.xMin) * (c.yMax + 1 - c.yMin) * (c.zMax + 1 - c.zMin)
}

func (c *cuboids) count() (count int) {
	for i := len(*c) - 1; i >= 0; i-- { // has to go backwards?
		cube := (*c)[i]
		if cube.action == "off" {
			continue
		}
		intersects, c1 := cuboids{}, &cube
		for _, c2 := range (*c)[i+1:] {
			intersect := cuboid{}
			intersect.xMin, intersect.xMax = math.IntMax(c1.xMin, c2.xMin), math.IntMin(c1.xMax, c2.xMax)
			intersect.yMin, intersect.yMax = math.IntMax(c1.yMin, c2.yMin), math.IntMin(c1.yMax, c2.yMax)
			intersect.zMin, intersect.zMax = math.IntMax(c1.zMin, c2.zMin), math.IntMin(c1.zMax, c2.zMax)
			if intersect.xMax < intersect.xMin || intersect.yMax < intersect.yMin || intersect.zMax < intersect.zMin {
				continue
			}
			intersect.volume = intersect.getVolume()
			intersect.action = "on"
			intersects = append(intersects, intersect)
		}
		count += cube.volume
		count -= intersects.count()
	}
	return count
}
