package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2015/input/1")
	in := strings.TrimSpace(string(input))
	floor, firstBasementVisit := 0, 0
	for i, c := range in {
		switch c {
		case '(': // up
			floor++
		case ')': // down
			floor--
		default:
			continue
		}
		if floor < 0 && firstBasementVisit == 0 {
			// slice starts at 0 and instructions state first position is 1
			firstBasementVisit = i + 1
		}
	}
	fmt.Println("Part 1:", floor)
	fmt.Println("Part 2:", firstBasementVisit)
}
