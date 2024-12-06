package day5

import (
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const DEBUG = false
const EXAMPLE = false

func Puzzle(input *[]byte, part2 bool) int {
	if EXAMPLE {
		*input, _ = os.ReadFile("2024/input/5.ex")
	}
	if DEBUG {
		log.SetLevel(log.DebugLevel)
	}

	// split into two sections on "\n\n"
	section := strings.Split(strings.TrimSpace(string(*input)), "\n\n")

	// parse each line of rules section
	rules := make(map[int]map[int]bool)
	for _, line := range strings.Split(section[0], "\n") {
		ruleStr := strings.Split(line, "|")
		b, _ := strconv.Atoi(ruleStr[0]) // before
		a, _ := strconv.Atoi(ruleStr[1]) // after
		if _, ok := rules[b]; ok {
			rules[b][a] = true
		} else {
			rules[b] = make(map[int]bool)
			rules[b][a] = true // a after b
		}
	}

	// parse each line of the updates section
	var updates [][]int
	for _, line := range strings.Split(section[1], "\n") {
		pages := strings.Split(line, ",")
		var update []int
		for _, page := range pages {
			pageNum, _ := strconv.Atoi(page)
			update = append(update, pageNum)
		}
		updates = append(updates, update)
	}

	// find all valid & invalid updates
	var validUpd [][]int
	var invalidUpd [][]int
	for _, update := range updates {
		if isValid(&update, &rules) {
			validUpd = append(validUpd, update)
		} else {
			invalidUpd = append(invalidUpd, update)
		}
	}

	if part2 {
		validUpd = [][]int{}
		log.Debug("This should be empty: ", validUpd)
		for _, update := range invalidUpd {
			for !isValid(&update, &rules) {
				bubbleSort(&update, &rules)
			}
			validUpd = append(validUpd, update)
		}
	}

	// total middle page numbers for valid
	totalValid := 0
	for _, update := range validUpd {
		totalValid += middle(&update)
	}
	return totalValid
}

func isValid(u *[]int, r *map[int]map[int]bool) bool {
	for i := 0; i < len(*u)-1; i++ {
		for j := i; j < len(*u); j++ {
			curr, next := (*u)[i], (*u)[j]
			if (*r)[next][curr] {
				return false
			}
		}
	}
	return true
}

func middle(u *[]int) int {
	return (*u)[(len(*u) / 2)]
}

func bubbleSort(u *[]int, r *map[int]map[int]bool) {
	for i := 0; i < len(*u)-1; i++ {
		for j := i + 1; j < len(*u); j++ {
			curr, next := (*u)[i], (*u)[j]
			if (*r)[next][curr] {
				(*u)[i], (*u)[j] = (*u)[j], (*u)[i] // swap
			}
		}
	}
}
