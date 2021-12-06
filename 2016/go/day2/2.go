package day2

import (
	"io/ioutil"
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
)

var filename = "2016/input/2"

func Part1() string {
	return Puzzle(Input(filename), false)
}

func Part2() string {
	return Puzzle(Input(filename), true)
}

func Input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func Puzzle(input []string, part2 bool) string {
	var code string
	var realCode string
	for _, instruction := range input {
		code = code + getButton(&instruction, 1)
		realCode = realCode + getButton(&instruction, 2)
	}
	if part2 {
		return realCode
	}
	return code
}

func getButton(instruction *string, keypadVersion int) string {
	keypad1 := make(map[geom.Point]string, 9)
	keypad1[geom.Point{-1, 1}] = "1"
	keypad1[geom.Point{0, 1}] = "2"
	keypad1[geom.Point{1, 1}] = "3"
	keypad1[geom.Point{-1, 0}] = "4"
	keypad1[geom.Point{0, 0}] = "5"
	keypad1[geom.Point{1, 0}] = "6"
	keypad1[geom.Point{-1, -1}] = "7"
	keypad1[geom.Point{0, -1}] = "8"
	keypad1[geom.Point{1, -1}] = "9"

	keypad2 := make(map[geom.Point]string, 9)
	keypad2[geom.Point{0, 2}] = "1"
	keypad2[geom.Point{-1, 1}] = "2"
	keypad2[geom.Point{0, 1}] = "3"
	keypad2[geom.Point{1, 1}] = "4"
	keypad2[geom.Point{-2, 0}] = "5"
	keypad2[geom.Point{-1, 0}] = "6"
	keypad2[geom.Point{0, 0}] = "7"
	keypad2[geom.Point{1, 0}] = "8"
	keypad2[geom.Point{2, 0}] = "9"
	keypad2[geom.Point{-1, -1}] = "A"
	keypad2[geom.Point{0, -1}] = "B"
	keypad2[geom.Point{1, -1}] = "C"
	keypad2[geom.Point{0, -2}] = "D"

	var keypad *map[geom.Point]string
	switch keypadVersion {
	case 1:
		keypad = &keypad1
	case 2:
		keypad = &keypad2
	}

	button := geom.Point{0, 0}
	for _, direction := range *instruction {
		tmp := button
		switch direction {
		case 'U':
			tmp.Y++
			if _, exists := (*keypad)[tmp]; exists {
				button.Y++
			}
		case 'D':
			tmp.Y--
			if _, exists := (*keypad)[tmp]; exists {
				button.Y--
			}
		case 'L':
			tmp.X--
			if _, exists := (*keypad)[tmp]; exists {
				button.X--
			}
		case 'R':
			tmp.X++
			if _, exists := (*keypad)[tmp]; exists {
				button.X++
			}
		}
	}
	return (*keypad)[button]
}
