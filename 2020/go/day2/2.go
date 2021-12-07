package day2

import (
	"fmt"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	p1, p2 := 0, 0
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		var min, max int
		var char byte
		var pw string
		fmt.Sscanf(i, "%v-%v %c: %v", &min, &max, &char, &pw)
		count := strings.Count(pw, string(char))
		if count >= min && count <= max {
			p1++
		}
		if (pw[min-1] == char) != (pw[max-1] == char) {
			p2++
		}
	}
	if part2 {
		return p2
	}
	return p1
}
