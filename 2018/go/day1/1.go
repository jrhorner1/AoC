package day1

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	frequency := 0
	changeList := []int{}
	for _, change := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
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
