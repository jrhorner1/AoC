package day21

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Monkey struct {
	id, lstr, rstr string
	value          uint64
	op             rune
	left, right    *Monkey
}

func Puzzle(input *[]byte, part2 bool) uint64 {
	monkeys := parseRiddle(input)
	if part2 {
		monkeys["humn"].value = 0
		for monkeys["root"].left.value == 0 && monkeys["root"].right.value == 0 {
			for i, monkey := range monkeys {
				if monkey.id != "humn" && monkey.value == 0 {
					if monkey.left.value != 0 && monkey.right.value != 0 {
						monkeys[i].number(false)
					}
				}
			}
		}
		if monkeys["root"].right.value != 0 {
			monkeys["root"].left.value = monkeys["root"].right.value
			monkeys["root"].left.number(part2)
		} else {
			monkeys["root"].right.value = monkeys["root"].left.value
			monkeys["root"].right.number(part2)
		}
		return monkeys["humn"].value
	}
	return monkeys["root"].number(part2)
}

func (m *Monkey) number(part2 bool) uint64 {
	a, b := m.left, m.right
	if part2 {
		if m.left.value > 0 && m.right.value == 0 {
			a, b = m.right, m.left
		}
		logrus.Debugf("Calculating %v|a %v:%d b %v:%d", m.id, a.id, a.value, b.id, b.value)
		switch m.op {
		case '+':
			a.value = m.value - b.value
			logrus.Debugf("%v:%d = %v:%d - %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
		case '-':
			if a == m.right {
				a.value = b.value - m.value
				logrus.Debugf("%v:%d = %v:%d - %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
			} else {
				a.value = m.value + b.value
				logrus.Debugf("%v:%d = %v:%d + %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
			}
		case '*':
			a.value = m.value / b.value
			logrus.Debugf("%v:%d = %v:%d / %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
		case '/':
			if a == m.right {
				a.value = b.value / m.value
				logrus.Debugf("%v:%d = %v:%d / %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
			} else {
				a.value = m.value * b.value
				logrus.Debugf("%v:%d = %v:%d * %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
			}
		default:
			logrus.Errorf("Unknown operation: {%c}", rune(m.op))
		}
		if a.id != "humn" || (a.left != nil && a.right != nil) {
			a.number(part2)
		}
	} else {
		if m.value > 0 {
			return m.value
		}
		if a.value == 0 {
			a.number(part2)
		}
		if b.value == 0 {
			b.number(part2)
		}
		switch m.op {
		case '+':
			m.value = a.value + b.value
			logrus.Debugf("%v:%d = %v:%d + %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
		case '-':
			m.value = a.value - b.value
			logrus.Debugf("%v:%d = %v:%d - %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
		case '*':
			m.value = a.value * b.value
			logrus.Debugf("%v:%d = %v:%d * %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
		case '/':
			m.value = a.value / b.value
			logrus.Debugf("%v:%d = %v:%d / %v:%d", a.id, a.value, m.id, m.value, b.id, b.value)
		default:
			logrus.Errorf("Unknown operation: {%c}", rune(m.op))

		}
	}
	return m.value
}

func parseRiddle(input *[]byte) map[string]*Monkey {
	monkeys := make(map[string]*Monkey)
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		monkey := Monkey{value: 0}
		m := strings.Split(strings.TrimSpace(line), ":")
		monkey.id = m[0]
		l := strings.Split(strings.TrimSpace(m[1]), " ")
		switch len(l) {
		case 1:
			number, err := strconv.Atoi(l[0])
			if err != nil {
				logrus.Error(err)
			}
			monkey.value = uint64(number)
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
		monkeys[monkey.id] = &monkey
	}
	for i, monkey := range monkeys {
		if monkey.lstr != "" && monkey.rstr != "" {
			monkeys[i].left = monkeys[monkey.lstr]
			monkeys[i].right = monkeys[monkey.rstr]
		}
	}
	return monkeys
}
