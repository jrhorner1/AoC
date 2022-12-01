package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	elves := []int{}
	for _, elf := range strings.Split(strings.TrimSpace(string(*input)), "\n\n") {
		sum := 0
		for _, calories_string := range strings.Split((strings.TrimSpace(elf)), "\n") {
			calories, err := strconv.Atoi(calories_string)
			if err != nil {
				fmt.Println(err)
			}
			sum += calories
		}
		elves = append(elves, sum)
	}
	sort.Ints(elves)
	mostCalories := elves[len(elves)-1]
	if part2 {
		top3 := elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
		return top3
	}
	return mostCalories
}
