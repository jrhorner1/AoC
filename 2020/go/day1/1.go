package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2020/input/1"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func puzzle(input []string, part2 bool) int {
	var entries []int
	for i := 0; i < len(input); i++ {
		tmp, _ := strconv.Atoi(input[i])
		entries = append(entries, tmp)
	}
	for i, x := range entries {
		for j, y := range entries[i+1:] {
			if part2 {
				for _, z := range entries[j+1:] {
					if x+y+z == 2020 {
						return x * y * z
					}
				}
			} else {
				if x+y == 2020 {
					return x * y
				}
			}
		}
	}
	return 0
}
