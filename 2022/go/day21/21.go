package day21

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Monkey struct {
	id, lstr, rstr string
	value          int
	op             rune
}

func Puzzle(input *[]byte, part2 bool) int {
	monkeys, lookup := parseRiddle(input)
	if part2 {
		return monkeys[lookup["humn"]].value
	}
	return valueOf("root", &monkeys, &lookup)
}

func valueOf(id string, monkeys *[]Monkey, lookup *map[string]int) int {
	monkey := (*monkeys)[(*lookup)[id]]
	if monkey.value != 0 {
		return monkey.value
	} else {
		switch monkey.op {
		case '+':
			(*monkeys)[(*lookup)[id]].value = valueOf(monkey.lstr, monkeys, lookup) + valueOf(monkey.rstr, monkeys, lookup)
		case '-':
			(*monkeys)[(*lookup)[id]].value = valueOf(monkey.lstr, monkeys, lookup) - valueOf(monkey.rstr, monkeys, lookup)
		case '*':
			(*monkeys)[(*lookup)[id]].value = valueOf(monkey.lstr, monkeys, lookup) * valueOf(monkey.rstr, monkeys, lookup)
		case '/':
			(*monkeys)[(*lookup)[id]].value = valueOf(monkey.lstr, monkeys, lookup) / valueOf(monkey.rstr, monkeys, lookup)
		default:
			logrus.Errorf("Unknown operation: %v", monkey.op)
		}
	}
	return (*monkeys)[(*lookup)[id]].value
}

func parseRiddle(input *[]byte) ([]Monkey, map[string]int) {
	lookup := make(map[string]int)
	monkeys := []Monkey{}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		monkey := Monkey{}
		m := strings.Split(strings.TrimSpace(line), ":")
		monkey.id = m[0]
		l := strings.Split(strings.TrimSpace(m[1]), " ")
		switch len(l) {
		case 1:
			number, err := strconv.Atoi(l[0])
			if err != nil {
				logrus.Error(err)
			}
			monkey.value = number
		case 3:
			monkey.lstr = l[0]
			monkey.rstr = l[2]
			if len(l[1]) != 1 {
				logrus.Errorf("Something went wrong: {%v}", l[1])
			}
			monkey.op = rune(l[1][0])
		default:
			logrus.Errorf("Something went wrong: {%v}", l)
		}
		monkeys = append(monkeys, monkey)
		lookup[monkey.id] = len(monkeys) - 1
	}
	return monkeys, lookup
}
