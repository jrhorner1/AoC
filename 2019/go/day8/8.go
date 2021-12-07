package day8

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var filename = "2019/input/8"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() {
	puzzle(input(filename), true)
}

func input(file string) string {
	input, _ := ioutil.ReadFile(file)
	output := strings.TrimSpace(string(input))
	return output
}

const (
	x int = 25
	y int = 6
)

func puzzle(input string, part2 bool) int {
	var layers [][]int

	adv := x * y
	for i := 0; i < len(input); i += adv {
		var layer []int
		layerString := input[i : i+adv]
		for _, r := range layerString {
			layer = append(layer, int(r)-48)
		}
		layers = append(layers, layer)
	}
	image := make([]int, x*y)
	for i := range image {
		image[i] = 2
	}
	var target, zctemp int = 0, x * y
	for i, layer := range layers {
		zcount := 0
		for j := range layer {
			if layer[j] == 0 {
				zcount++
			}
			if image[j] == 2 {
				switch layer[j] {
				case 0: // black
					image[j] = layer[j]
				case 1: // white
					image[j] = layer[j]
				case 2: // transparent
					continue
				}
			}
		}
		if zcount < zctemp {
			target = i
			zctemp = zcount
		}
	}
	var one, two int = 0, 0
	for i := range layers[target] {
		switch layers[target][i] {
		case 1:
			one++
		case 2:
			two++
		default:
			continue
		}
	}
	if part2 {
		for i := range image {
			for j := 1; j <= y; j++ {
				if i == x*j {
					fmt.Println()
				}
			}
			if image[i] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(image[i])
			}
		}
	}
	return one * two
}
