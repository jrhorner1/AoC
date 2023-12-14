package day12

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	records, sum := []int{}, 0
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		format := strings.Split(strings.TrimSpace(line), " ")
		record := format[0]
		groups := []int{}
		for _, numStr := range strings.Split(format[1], ",") {
			numInt, err := strconv.Atoi(numStr)
			if err != nil {
				log.Error(err)
			}
			groups = append(groups, numInt)
		}
		if part2 {
			r, g := record, groups
			for i := 1; i < 5; i++ {
				record += "?" + r
				groups = append(groups, g...)
			}
		}
		log.Debugf("Record: %s Groups: %d", record, groups)
		permutations := permutate(record, groups)
		records = append(records, permutations)
		sum += permutations
	}
	log.Debug(records, sum)
	return sum
}

func permutate(record string, groups []int) int {
	if len(groups) == 0 {
		if strings.Contains(record, "#") {
			return 0
		}
		return 1
	}
	if len(record) == 0 {
		return 0
	}
	if len(record) < groups[0] {
		return 0
	}
	// log.Debugf("Checking %c, group size %d", rune(record[0]), groups[0])
	switch rune(record[0]) {
	case '.': // working
		return permutate(record[1:], groups)
	case '#': // broken
		return possiblites(record, groups)
	case '?': // unknown
		return permutate(record[1:], groups) + possiblites(record, groups)
	default:
		return 0
	}
}

func possiblites(r string, g []int) int {
	s := g[0]
	if len(r) < s {
		// log.Debugf("Record length less than group size. %s %d", r, g)
		return 0
	}
	sub := r[:s]
	sub = strings.ReplaceAll(sub, "?", "#")
	if sub != strings.Repeat("#", s) {
		// log.Debugf("Substring can't fit the group. %s %s %d", r, sub, g)
		return 0
	}
	if len(r) == s {
		if len(g) != 1 {
			// log.Debugf("Groups leftover. %s %s %d", r, sub, g)
			return 0
		}
		return 1
	}
	if r[s] == '?' || r[s] == '.' {
		log.Debug("Current: ", r, g)
		log.Debug("Next: ", r[s+1:], g[1:])
		return permutate(r[s+1:], g[1:])
	}
	// log.Debugf("No possilities. %s %s %d", r, sub, g)
	return 0
}
