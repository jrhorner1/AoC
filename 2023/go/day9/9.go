package day9

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	values := []int{}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		fields := strings.Fields(line)
		history := []int{}
		for _, field := range fields {
			number, err := strconv.Atoi(field)
			if err != nil {
				log.Error(err)
			}
			history = append(history, number)
		}
		next := extrapolate(&history, &part2)
		log.Debug("Next: ", next)
		values = append(values, next)
	}
	log.Debug("Values: ", values)
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func extrapolate(history *[]int, part2 *bool) int {
	first := (*history)[0]
	last := (*history)[len(*history)-1]
	increment := 0
	diffs := []int{}
	for i := range *history {
		if i == 0 {
			continue
		}
		diffs = append(diffs, (*history)[i]-(*history)[i-1])
	}
	log.Debug("Diffs: ", diffs)
	if allSame(&diffs) {
		increment = diffs[0]
	} else {
		increment = extrapolate(&diffs, part2)
	}
	log.Debug("Increment: ", increment)
	if *part2 {
		return first - increment
	}
	return last + increment
}

func allSame(diffs *[]int) bool {
	first := (*diffs)[0]
	for _, diff := range (*diffs)[1:] {
		if diff != first {
			return false
		}
	}
	return true
}
