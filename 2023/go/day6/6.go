package day6

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	log.Debug("Insert Code Here")
	lines := strings.Split(strings.TrimSpace(string(*input)), "\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	margin := 1
	for i := range times {
		if i == 0 {
			continue
		}
		margin *= simulateRace(times[i], distances[i])
	}
	if part2 {
		time := strings.Split(strings.ReplaceAll(lines[0], " ", ""), ":")
		distance := strings.Split(strings.ReplaceAll(lines[1], " ", ""), ":")
		return simulateRace(time[1], distance[1])
	}
	return margin
}

func simulateRace(t, d string) int {
	time, _ := strconv.Atoi(t)
	distance, _ := strconv.Atoi(d)
	waysToWin := 0
	for hold := 0; hold <= time; hold++ {
		speed := hold
		travel := speed * (time - hold)
		if travel > distance {
			waysToWin++
		}
	}
	return waysToWin
}
