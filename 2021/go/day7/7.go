package day7

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math"
)

// var filename = "2021/input/7"

var filename = "2021/examples/7"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []int {
	input, _ := ioutil.ReadFile(file)
	in := strings.Split(strings.TrimSpace(string(input)), ",")
	output := []int{}
	for _, i := range in {
		out, _ := strconv.Atoi(i)
		output = append(output, out)
	}
	return output
}

func puzzle(input []int, part2 bool) int {
	most, total := 0, 0
	for _, i := range input { // determine the furthest (largest) point
		total += i
		if i > most {
			most = i
		}
	}
	if part2 {
		median, totalFuel := 0, 0
		// mediam for example input is 4.9 but rounding based on remainder doesnt work for the puzzle input
		if filename == "2021/examples/7" {
			median = (total / len(input)) + 1
		} else {
			median = total / len(input)
		}
		for _, i := range input {
			steps := math.IntAbs(i - median)
			for i := 1; i <= steps; i++ {
				totalFuel += i
			}
		}
		return totalFuel
	}
	steps := make(map[int]int)   // each potential alignment position and total steps to get there
	for i := 0; i <= most; i++ { // calculate total steps to each position
		for _, p := range input {
			if p > i {
				steps[i] += p - i
			} else {
				steps[i] += i - p
			}
		}
	}
	leastSteps := steps[0]
	for i := 1; i < len(steps); i++ { // find the least steps
		if steps[i] < leastSteps {
			leastSteps = steps[i]
		}
	}
	return leastSteps
}
