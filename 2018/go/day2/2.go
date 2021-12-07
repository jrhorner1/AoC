package day2

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n")
	count2, count3 := 0, 0
	for _, id := range in {
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
		for i, idA := range in {
			for j, idB := range in {
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
