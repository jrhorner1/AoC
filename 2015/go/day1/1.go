package day1

import (
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	inputString := strings.TrimSpace(string(*input))
	floor, firstBasementVisit := 0, 0
	for i, c := range inputString {
		switch c {
		case '(': // up
			floor++
		case ')': // down
			floor--
		default:
			continue
		}
		if part2 && floor < 0 && firstBasementVisit == 0 {
			// slice starts at 0 and instructions state first position is 1
			return i + 1
		}
	}
	return floor
}
