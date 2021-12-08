package day8

import (
	"sort"
	"strconv"
	"strings"
)

/*
segments: digit
7: "8"
6: "0,6,9"
5: "2,3,5"
4: "4"
3: "7"
2: "1"
*/

func Puzzle(input *[]byte, part2 bool) int {
	// parse the input
	wires, displays := [][]string{}, [][]string{}
	uniqSegments := map[int]string{2: "1", 4: "4", 3: "7", 7: "8"}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") { // split on newline
		signals, display := []string{}, []string{}
		for i, section := range strings.Split(strings.TrimSpace(line), "|") { // split on |
			switch i {
			case 0: // parse the signals
				for _, s := range strings.Split(strings.TrimSpace(section), " ") {
					signal := []rune(s)
					sortRunes(signal) // sort characters so they can be matched with the digit
					signals = append(signals, string(signal))
				}
			case 1:
				for _, d := range strings.Split(strings.TrimSpace(section), " ") {
					digit := []rune(d)
					sortRunes(digit) // sort characters to be matched against signal
					display = append(display, string(digit))
				}
			}
		}
		sortSignals(signals) // sort signals by their length, putting them in the correct order to determine each digits display signal
		wires = append(wires, signals)
		displays = append(displays, display)
	}
	if part2 {
		outputTotal := 0
		for i, signals := range wires {
			segments, stnemges, digits := make(map[rune]int), make(map[int]rune), make(map[string]string)
			rightSide, topLeftCenter := "", ""
			for _, signal := range signals {
				switch len(signal) {
				case 2:
					digits[signal] = "1"
					rightSide = signal
				case 3:
					digits[signal] = "7"
					for _, r := range signal {
						if !strings.ContainsRune(rightSide, r) {
							segments[r] = 0 // top
						}
					}
				case 4:
					digits[signal] = "4"
					for _, r := range signal {
						if !strings.ContainsRune(rightSide, r) {
							topLeftCenter += string(r)
						}
					}
				case 5:
					if strings.ContainsRune(signal, rune(rightSide[0])) && strings.ContainsRune(signal, rune(rightSide[1])) {
						digits[signal] = "3"
						for _, r := range topLeftCenter {
							if strings.ContainsRune(signal, r) {
								segments[r] = 3 // center
							} else {
								segments[r] = 1 // top left
							}
						}
					} else {
						if strings.ContainsRune(signal, rune(topLeftCenter[0])) && strings.ContainsRune(signal, rune(topLeftCenter[1])) {
							digits[signal] = "5"
							for _, r := range rightSide {
								if strings.ContainsRune(signal, r) {
									segments[r] = 5 // bottom right
								} else {
									segments[r] = 2 // top right
								}
							}
							for _, r := range signal {
								if _, found := segments[r]; !found {
									segments[r] = 6 // bottom
								}
							}
						} else { // don't write any segments from 2 as its ambiguous until segment 6 (bottom) is determined
							digits[signal] = "2"
						}
					}
					if len(segments) == 6 { // segment 4 should be the only one left
						for _, r := range []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'} {
							if _, found := segments[r]; !found {
								segments[r] = 4 // bottom left
							}
						}
					}
				case 6: // use the completed segments map to finish digits map
					if !strings.ContainsRune(signal, stnemges[3]) {
						digits[signal] = "0"
					} else if !strings.ContainsRune(signal, stnemges[2]) {
						digits[signal] = "6"
					} else if !strings.ContainsRune(signal, stnemges[4]) {
						digits[signal] = "9"
					}
				case 7:
					digits[signal] = "8"
				}
				if len(segments) == 7 { // if the segments map is complete
					for k, v := range segments {
						stnemges[v] = k // invert it for searching display output
					}
				}
			}
			outputString := ""
			for _, display := range displays[i] {
				outputString += digits[string(display)]
			}
			output, _ := strconv.Atoi(outputString)
			outputTotal += output
		}
		return outputTotal
	}
	countUniqSegs := 0
	for _, display := range displays {
		for _, digit := range display {
			if _, ok := uniqSegments[len(digit)]; ok {
				countUniqSegs += 1
			}
		}
	}
	return countUniqSegs
}

func sortSignals(input []string) {
	sort.Sort(StringSlice(input))
}

type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return len(x[i]) < len(x[j]) }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func sortRunes(input []rune) {
	sort.Sort(RuneSlice(input))
}

type RuneSlice []rune

func (x RuneSlice) Len() int           { return len(x) }
func (x RuneSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x RuneSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
