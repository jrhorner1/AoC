package day6

import (
	"regexp"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	// countAny, countEvery := 0, 0
	count := 0
	if part2 {
		count = 0
		for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n\n") {
			split := strings.Split(i, "\n")
			q := map[rune]bool{}
			for i := 0; i < 26; i++ {
				q[rune(int('a')+i)] = true
			}
			for _, j := range split {
				for k, _ := range q {
					match, _ := regexp.MatchString(string(k), j)
					if !match {
						q[k] = match
					}
				}
			}
			for _, v := range q {
				if v {
					count++
				}
			}
		}
	} else {
		q := "abcdefghijklmnopqrstuvwxyz"
		for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n\n") {
			for _, j := range q {
				match, _ := regexp.MatchString(string(j), i)
				if match {
					count++
				}
			}
		}
	}
	return count
}
