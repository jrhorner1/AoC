package day2

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2018/input/2"

func Part1() int {
	return Puzzle(Input(filename), false)
}

func Part2() int {
	return Puzzle(Input(filename), true)
}

func Input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func Puzzle(input []string, part2 bool) int {
	count2, count3 := 0, 0
	for _, id := range input {
		chars := make(map[rune]int)
		for _, char := range id {
			chars[char] += 1
		}
		c2, c3 := false, false
		for _, v := range chars {
			switch v {
			case 2:
				c2 = true
			case 3:
				c3 = true
			}
		}
		if c2 {
			count2++
		}
		if c3 {
			count3++
		}
	}
	if part2 {
		for i, idA := range input {
			for j, idB := range input {
				if i == j {
					continue
				}
				diff := []int{}
				for i := 0; i < len(idA); i++ {
					if idA[i] != idB[i] {
						diff = append(diff, i)
					}
				}
				if len(diff) == 1 {
					out, _ := strconv.Atoi(idA[:diff[0]] + idB[diff[0]+1:])
					return out
				}
			}
		}
	}
	return count2 * count3
}
