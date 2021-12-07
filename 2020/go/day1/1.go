package day1

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n")
	var entries []int
	for i := 0; i < len(in); i++ {
		tmp, _ := strconv.Atoi(in[i])
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
