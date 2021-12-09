package day8

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	// parse the input
	wires, displays := [][]string{}, [][]string{}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") { // split on newline
		signals, display := []string{}, []string{}
		for i, section := range strings.Split(strings.TrimSpace(line), "|") { // split on |
			switch i {
			case 0: // parse the signals
				for _, signal := range strings.Split(strings.TrimSpace(section), " ") {
					signals = append(signals, signal)
				}
			case 1:
				for _, digit := range strings.Split(strings.TrimSpace(section), " ") {
					display = append(display, digit)
				}
			}
		}
		wires = append(wires, signals)
		displays = append(displays, display)
	}
	if part2 {
		outputTotal := 0
		sums := map[int]string{42: "0", 17: "1", 34: "2", 39: "3", 30: "4", 37: "5", 41: "6", 25: "7", 49: "8", 45: "9"}
		for i, signals := range wires {
			score := make(map[rune]int)
			for _, signal := range signals {
				for _, r := range signal {
					score[r] += 1
				}
			}
			outputString := ""
			for _, display := range displays[i] {
				sum := 0
				for _, r := range display {
					sum += score[r]
				}
				outputString += sums[sum]
			}
			output, _ := strconv.Atoi(outputString)
			outputTotal += output
		}
		return outputTotal
	}
	uniqueSignals := map[int]string{2: "1", 4: "4", 3: "7", 7: "8"}
	countUniqueSignals := 0
	for _, display := range displays {
		for _, digit := range display {
			if _, found := uniqueSignals[len(digit)]; found {
				countUniqueSignals += 1
			}
		}
	}
	return countUniqueSignals
}
