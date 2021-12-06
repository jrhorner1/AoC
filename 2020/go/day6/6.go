package day6

import (
	"io/ioutil"
	"regexp"
	"strings"
)

var filename = "2020/input/6"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	return output
}

func puzzle(input []string, part2 bool) int {
	// countAny, countEvery := 0, 0
	count := 0
	q := "abcdefghijklmnopqrstuvwxyz"
	for _, i := range input {
		for _, j := range q {
			match, _ := regexp.MatchString(string(j), i)
			if match {
				count++
			}
		}
	}
	if part2 {
		count = 0
		for _, i := range input {
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
	}
	return count
}
