package day10

import (
	"sort"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	jolts := []int{0}
	diffs := map[int]int{}
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		adapter, _ := strconv.Atoi(i)
		jolts = append(jolts, adapter)
	}
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	arrs := map[int]int{0: 1}
	for i := 1; i < len(jolts); i++ {
		diffs[jolts[i]-jolts[i-1]]++
		for j := 0; j < i; j++ {
			if (jolts[i] - jolts[j]) <= 3 {
				arrs[i] += arrs[j]
			}
		}
	}
	if part2 {
		return arrs[len(jolts)-1]
	}
	return diffs[1] * diffs[3]
}
