package day1

import (
	"regexp"
	"strconv"
	"strings"

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
	kv := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for k, v := range kv {
		re := regexp.MustCompile(k)
		cs = string(re.ReplaceAll([]byte(cs), []byte(v)))
	}
	return cs
}
