package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2020/input/10")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	jolts := []int{0}
	diffs := map[int]int{}
	for _, i := range in {
		adapter, _ := strconv.Atoi(i)
		jolts = append(jolts, adapter)
	}
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	arrs := map[int]int64{0: 1}
	for i := 1; i < len(jolts); i++ {
		diffs[jolts[i]-jolts[i-1]]++
		for j := 0; j < i; j++ {
			if (jolts[i] - jolts[j]) <= 3 {
				arrs[i] += arrs[j]
			}
		}
	}
	fmt.Println("Part 1:", diffs[1]*diffs[3])
	fmt.Println("Part 2:", arrs[len(jolts)-1])
}
