package day4

import (
	"regexp"
	"strings"
)

var regs = []*regexp.Regexp{
	regexp.MustCompile(`(?:^|\s)(byr):(?:(19[2-9]\d|200[0-2])(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(iyr):(?:(201\d|2020)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(eyr):(?:(202\d|2030)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(hgt):(?:((?:1[5-8]\d|19[0-3])cm|(?:59|6\d|7[0-6])in)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(hcl):(?:(#[\da-f]{6})(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(ecl):(?:(amb|blu|brn|gry|grn|hzl|oth)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(pid):(?:(\d{9})(?:\s|$))?`)}

func Puzzle(input *[]byte, part2 bool) int {
	p1, p2 := 0, 0
	for _, pp := range strings.Split(strings.TrimSpace(string(*input)), "\n\n") {
		d1, d2 := 1, 1
		for _, reg := range regs {
			if m := reg.FindStringSubmatch(pp); len(m) == 0 {
				d1, d2 = 0, 0
			} else if m[2] == "" {
				d2 = 0
			}
		}
		p1, p2 = p1+d1, p2+d2
	}
	if part2 {
		return p2
	}
	return p1
}
