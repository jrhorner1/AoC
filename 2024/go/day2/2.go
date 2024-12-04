package day2

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	// log.SetLevel(log.DebugLevel)
	// safe report counter
	safe := 0
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		// parse the report
		var report []int
		for _, str := range strings.Split(strings.TrimSpace((line)), " ") {
			value, _ := strconv.Atoi(str)
			report = append(report, value)
		}
		// check if safe
		if isSafe(&report, part2) {
			safe++
		}
	}
	return safe
}

func isSafe(report *[]int, part2 bool) bool {
	log.Debug(*report)

	// determine asc or desc
	desc := true
	front, rear := (*report)[0], (*report)[len(*report)-1]
	if front-rear < 0 {
		desc = false
	}

	safe := true
	for i, num := range *report {
		// make sure i+1 exists
		if i+1 < len(*report) {
			// get a positive diff based on asc/desc
			diff := (*report)[i+1] - num
			if desc {
				diff = num - (*report)[i+1]
			}
			// check if not 1, 2, or 3
			if diff < 1 || diff > 3 {
				safe = false
			}
			log.Debug("safe: ", safe, ", diff: ", diff)
		}
	}

	if !safe && part2 {
		for i := range *report {
			// remove element i and recheck if safe
			temp := make([]int, len(*report))
			copy(temp, *report)
			if i+1 < len(*report) {
				temp = append(temp[:i], temp[i+1:]...)
			} else {
				temp = temp[:i]
			}
			log.Debug("temp: ", temp)
			if safe = isSafe(&temp, false); safe {
				return safe
			}
		}
	}

	return safe
}
