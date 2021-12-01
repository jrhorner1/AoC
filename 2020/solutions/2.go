package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2020/input/2")
	p1, p2 := 0, 0
	for _, i := range strings.Split(strings.TrimSpace(string(input)), "\n") {
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
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
