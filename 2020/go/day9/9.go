package day9

import (
	"sort"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	xmas := []int{}
	for _, s := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		v, _ := strconv.Atoi(s)
		xmas = append(xmas, v)
	}
	valid := map[int]bool{}
	var invalid int
	for i, data := range xmas {
		if i < 25 {
			continue
		}
		preamble := xmas[i-25 : i]
		for j, x := range preamble {
			for _, y := range preamble[j:] {
				if x+y == data {
					valid[data] = true
				}
			}
		}
		if !valid[data] {
			invalid = data
			if part2 {
				break
			}
			return data
		}
	}
	for low, x := range xmas {
		total := x
		for high, y := range xmas {
			if high <= low {
				continue
			}
			total += y
			if total > invalid {
				break
			} else if total == invalid {
				sort.Ints(xmas[low : high+1])
				return xmas[low] + xmas[high]
			}
		}
	}
	return 0
}
