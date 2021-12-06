package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2018/input/1"

func Part1() int {
	return Puzzle(Input(filename), false)
}

func Part2() int {
	return Puzzle(Input(filename), true)
}

func Input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func Puzzle(input []string, part2 bool) int {
	frequency := 0
	changeList := []int{}
	for _, change := range input {
		direction := change[0]
		amount, _ := strconv.Atoi(change[1:])
		switch direction {
		case '+':
			frequency += amount
			changeList = append(changeList, amount)
		case '-':
			frequency -= amount
			changeList = append(changeList, -amount)
		}
	}
	if part2 {
		frequencyLog := []int{0}
		frequency = 0
	loop:
		for {
			for _, change := range changeList {
				frequency += change
				for _, log := range frequencyLog {
					if frequency == log {
						break loop
					}
				}
				frequencyLog = append(frequencyLog, frequency)
			}
		}
		return frequency
	}
	return frequency
}
