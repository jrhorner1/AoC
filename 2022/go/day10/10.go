package day10

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	width  = 40
	height = 6
)

var (
	registerX      = 1
	cycle          = 0
	signalStrength []int
	crt            []bool
	image          string
)

func Puzzle(input *[]byte) (int, string) {
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		instruction := strings.Split(strings.TrimSpace(line), " ")
		switch instruction[0] {
		case "noop":
			tick(&cycle, &registerX, &signalStrength, &crt)
		case "addx":
			tick(&cycle, &registerX, &signalStrength, &crt)
			tick(&cycle, &registerX, &signalStrength, &crt)
			value, err := strconv.Atoi(instruction[1])
			if err != nil {
				logrus.Error(err)
			}
			registerX += value
		}
	}
	sum := 0
	for _, sigStr := range signalStrength {
		sum += sigStr
	}
	for i, pixel := range crt {
		if pixel {
			image += "#"
		} else {
			image += "."
		}
		if (i+1)%width == 0 {
			image += "\n"
		}
	}
	return sum, image
}

func tick(cycle, registerX *int, signalStrength *[]int, crt *[]bool) {
	sprite := []int{*registerX - 1, *registerX, *registerX + 1}
	*crt = append(*crt, false)
	for _, pixel := range sprite {
		if pixel == *cycle%width {
			(*crt)[*cycle] = true
		}
	}
	*cycle++
	if (*cycle-20)%40 == 0 {
		*signalStrength = append(*signalStrength, *cycle**registerX)
	}
}
