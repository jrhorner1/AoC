package day1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	//log.SetLevel(log.DebugLevel)
	total := 0
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		log.Debug(line)
		calibrationString := line
		if part2 {
			calibrationString = parseDigit(line)
		}
		re := regexp.MustCompile("[0-9]")
		digits := re.FindAllString(calibrationString, -1)
		log.Debug(digits)
		var first, last string
		for _, digit := range digits {
			if first == "" {
				first = digit
			}
			last = digit
		}
		valueStr := first + last
		log.Debug(valueStr)
		calibrationValue, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Error(err)
		}
		total += calibrationValue
	}
	return total
}

func parseDigit(cs string) string {
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	ncs := ""
	for i, r := range cs {
		if unicode.IsDigit(r) {
			ncs += string(r)
			continue
		}
		for d, digit := range digits {
			re := regexp.MustCompile(digit)
			if re.Match([]byte(cs[i:])) {
				ncs += fmt.Sprint(d + 1)
			}
		}
	}
	return ncs
}
