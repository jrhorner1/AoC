package day3

import (
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	// log.SetLevel(log.DebugLevel)
	var valid []string
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
		if part2 {
			re = regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\))`)
		}
		instr := re.FindAllString(line, -1)
		valid = append(valid, instr...)
	}
	log.Debug(valid)

	total := 0
	en := true
	for _, instr := range valid {
		switch instr {
		case "do()":
			en = true
		case "don't()":
			en = false
		default:
			if en {
				re := regexp.MustCompile(`[0-9]{1,3},[0-9]{1,3}`)
				str := strings.Split(re.FindString(instr), ",")
				a, _ := strconv.Atoi(str[0])
				b, _ := strconv.Atoi(str[1])
				total += a * b
			}
		}
	}

	if part2 {
		return total
	}
	return total
}
