package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type point struct {
	x int // left/right
	y int // up/down
}

func main() {
	input, _ := ioutil.ReadFile("2016/input/2")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	var code string
	var realCode string
	for _, instruction := range in {
		code = code + getButton(&instruction, 1)
		realCode = realCode + getButton(&instruction, 2)
	}
	fmt.Println("Part 1:", code)
	fmt.Println("Part 2:", realCode)
}

func getButton(instruction *string, keypadVersion int) string {
	keypad1 := make(map[point]string, 9)
	keypad1[point{-1, 1}] = "1"
	keypad1[point{0, 1}] = "2"
	keypad1[point{1, 1}] = "3"
	keypad1[point{-1, 0}] = "4"
	keypad1[point{0, 0}] = "5"
	keypad1[point{1, 0}] = "6"
	keypad1[point{-1, -1}] = "7"
	keypad1[point{0, -1}] = "8"
	keypad1[point{1, -1}] = "9"

	keypad2 := make(map[point]string, 9)
	keypad2[point{0, 2}] = "1"
	keypad2[point{-1, 1}] = "2"
	keypad2[point{0, 1}] = "3"
	keypad2[point{1, 1}] = "4"
	keypad2[point{-2, 0}] = "5"
	keypad2[point{-1, 0}] = "6"
	keypad2[point{0, 0}] = "7"
	keypad2[point{1, 0}] = "8"
	keypad2[point{2, 0}] = "9"
	keypad2[point{-1, -1}] = "A"
	keypad2[point{0, -1}] = "B"
	keypad2[point{1, -1}] = "C"
	keypad2[point{0, -2}] = "D"

	var keypad *map[point]string
	switch keypadVersion {
	case 1:
		keypad = &keypad1
	case 2:
		keypad = &keypad2
	}

	button := point{0, 0}
	for _, direction := range *instruction {
		tmp := button
		switch direction {
		case 'U':
			tmp.y++
			if _, exists := (*keypad)[tmp]; exists {
				button.y++
			}
		case 'D':
			tmp.y--
			if _, exists := (*keypad)[tmp]; exists {
				button.y--
			}
		case 'L':
			tmp.x--
			if _, exists := (*keypad)[tmp]; exists {
				button.x--
			}
		case 'R':
			tmp.x++
			if _, exists := (*keypad)[tmp]; exists {
				button.x++
			}
		}
	}
	return (*keypad)[button]
}
