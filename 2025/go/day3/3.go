package day3

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
	joltage := 0
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		var bank []string
		bank = append(bank, strings.Split(strings.TrimSpace(line), "")...)
		log.Debug(bank)
		n := 1 // 2 batteries
		if part2 {
			n = 11 // 12 batteries
		}
		var battery []string
		for len(battery) != n+1 {
			battery = append(battery, "0")
		}
		index := -1
		for i := 0; i <= n; i++ {
			for j := index + 1; j < len(bank)-(n-i); j++ {
				if jolts := bank[j]; jolts > battery[i] {
					battery[i], index = jolts, j
					// log.Debug(battery[1])
				}
			}

		}
		log.Debug(battery)
		bankMaxJoltage, err := strconv.Atoi(strings.Join(battery, ""))
		// log.Debug(bankMaxJoltage)
		if err != nil {
			log.Error(err)
		}
		joltage += bankMaxJoltage
	}
	return joltage
}
