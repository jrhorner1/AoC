package day6

import (
	"math"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	//log.SetLevel(log.DebugLevel)
	lines := strings.Split(strings.TrimSpace(string(*input)), "\n")
	if part2 {
		time, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(lines[0], " ", ""), ":")[1])
		distance, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(lines[1], " ", ""), ":")[1])
		return quadriatic(float64(time), float64(distance))
	}
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	margin := 1
	for i := range times {
		if i == 0 {
			continue
		}
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		margin *= quadriatic(float64(time), float64(distance))
		log.Debug("Margin: ", margin)
	}
	return margin
}

func quadriatic(t, d float64) int {
	x := math.Sqrt(float64(t*t - 4.0*d))
	max := math.Ceil((t + x) / 2.0)
	if math.Mod(max, 1.0) == 0.0 {
		max--
	}
	min := math.Floor((t - x) / 2.0)
	if math.Mod(min, 1.0) == 0.0 {
		min++
	}
	log.Debug("QMin: ", min, " QMax: ", max)
	return int(math.Abs(max-min)) + 1
}
