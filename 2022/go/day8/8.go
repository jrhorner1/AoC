package day8

import (
	"image"
	"strings"
)

var forestSize = 0

func Puzzle(input *[]byte, part2 bool) int {
	forest := Forest{}
	for y, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		if y == 0 {
			forestSize = len(line)
		}
		for x, char := range line {
			position := image.Point{x, y}
			tree := Tree{height: int(char) - '0', isVisible: false, scenicScore: 1}
			forest[position] = tree
		}
	}
	if part2 {
		return forest.survey("idealTreehouse")
	}
	return forest.survey("visibleTrees")
}

type Tree struct {
	height,
	scenicScore int
	isVisible bool
}

type Forest map[image.Point]Tree

func (f *Forest) survey(arg string) int {
	numVisible, topScore := 0, 0
	for location, tree := range *f {
		for _, delta := range []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			viewingDistance := 0
			for position := location.Add(delta); ; position = position.Add(delta) {
				if compare, ok := (*f)[position]; !ok {
					tree.isVisible = true
					tree.scenicScore *= viewingDistance
					break
				} else if viewingDistance++; compare.height >= tree.height {
					tree.scenicScore *= viewingDistance
					break
				}
			}
		}
		if tree.isVisible {
			numVisible++
		}
		if tree.scenicScore > topScore {
			topScore = tree.scenicScore
		}
	}
	switch arg {
	case "visibleTrees":
		return numVisible
	case "idealTreehouse":
		return topScore
	default:
		return 0
	}
}
