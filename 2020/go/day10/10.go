package day10

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var filename = "2020/input/10"

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
	jolts := []int{0}
	diffs := map[int]int{}
	for _, i := range input {
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
