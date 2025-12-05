package day1

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const DEBUG = false

func Puzzle(input *[]byte, part2 bool) int {
	if DEBUG {
		log.SetLevel(log.DebugLevel)
	}
	dial := 50
	zeros := 0

	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		log.Debugf("Instruction: %s:", line)
		dir := rune(line[0])
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Error(err)
		}
		if part2 {
			passes := n / 100
			zeros += passes
			log.Debugf("Initial passes: %d", passes)
		}
		r := n % 100
		log.Debugf("Remainder: %d", r)
		switch dir {
		case 'R': // add r to dial
			if r >= (100 - dial) {
				dial = r - (100 - dial)
				if part2 && (dial != 0 && dial != 100) {
					zeros++
				}
			} else {
				dial += r
			}
		case 'L': // subtract r from dial
			if dial == 0 {
				dial = 100
			}
			if r >= dial {
				dial = 100 - (r - dial)
				if part2 && (dial != 0 && dial != 100) {
					zeros++
				}
			} else {
				dial -= r
			}
		default:
			log.Error("I don't know what direction that is, but it ain't left or right.")
		}
		if dial == 100 {
			dial = 0
		}
		log.Debugf("Dial stopped on: %d", dial)
		if dial == 0 {
			zeros++
		}
		log.Debugf("Zeros: %d", zeros)
	}
	return zeros
}
