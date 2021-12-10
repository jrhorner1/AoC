package day5

import (
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	nice := []string{}
santasList:
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		if part2 {
			// any pair of runes twice, but not overlapping ie. eyey, ddlskdd
			for i := 0; i < len(line)-2; i++ {
				substring := line[i : i+2]
				newLine := line[i+2:]
				if strings.Contains(newLine, substring) {
					// any single rune twice with any other single rune between, ie. aba, xyx, iii
					for i := 0; i < len(line)-2; i++ {
						if string(line[i]) == string(line[i+2]) {
							nice = append(nice, line)
							continue santasList
						}
					}
				}
			}
		} else {
			// not "ab", "cd", "pq", or "xy"
			naughtySubstrings := []string{"ab", "cd", "pq", "xy"}
			for _, substring := range naughtySubstrings {
				if strings.Contains(line, substring) {
					continue santasList
				}
			} // a double letter
			doubles := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"}
			for _, substring := range doubles {
				if strings.Contains(line, substring) {
					// 3 vowels
					vowels := []rune{'a', 'e', 'i', 'o', 'u'}
					vowelCount := 0
					for _, vowel := range vowels {
						for _, r := range line {
							if r == vowel {
								vowelCount++
							}
						}
					}
					if vowelCount >= 3 {
						nice = append(nice, line)
						continue santasList
					}
				}
			}
		}
	}
	return len(nice)
}
