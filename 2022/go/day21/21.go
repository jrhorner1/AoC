package day21

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	lines := make(map[string]string)
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		l := strings.Split(strings.TrimSpace(line), ":")
		lines[l[0]] = l[1]
	}
	monkeys := make(map[string]int)
	for len(lines) != 0 {
		for id, value := range lines {
			if len(value) < 11 {
				number, err := strconv.Atoi(strings.TrimSpace(value))
				if err != nil {
					logrus.Error(err)
				}
				monkeys[id] = number
				delete(lines, id)
			}
			if len(value) >= 11 {
				l := strings.Split(strings.TrimSpace(value), " ")
				if a, ok := monkeys[l[0]]; ok {
					if b, ok := monkeys[l[2]]; ok {
						switch l[1] {
						case "*":
							monkeys[id] = a * b
						case "/":
							monkeys[id] = a / b
						case "+":
							monkeys[id] = a + b
						case "-":
							monkeys[id] = a - b
						default:
							logrus.Errorf("What kind of operation is that? ( %s )", l[1])
						}
						delete(lines, id)
					}
				}
			}
		}
	}
	if part2 {
		return monkeys["humn"]
	}
	return monkeys["root"]
}
