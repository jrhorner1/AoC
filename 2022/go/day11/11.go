package day11

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	monkeys := parseMonkeys(input)
	rounds := 20
	tests := []int{}
	lcm := 0
	if part2 {
		rounds = 10000
		for _, monkey := range monkeys {
			tests = append(tests, int(monkey.test.value))
		}
		lcm = lowestCommonMultiple(tests[0], tests[1])
		for i := 2; i < len(tests); i++ {
			lcm = lowestCommonMultiple(lcm, tests[i])
		}
	}
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				// inspects item
				monkey.inspect(item)
				// gets bored,
				if part2 { // modular arithmatic to manage worry level
					*item = *item % lcm
				} else { // worry level divided by 3, rounded down
					*item = int(math.Floor(float64(*item) / 3))
				}
				// test operation
				if *item%monkey.test.value == 0 {
					monkeys[monkey.test.yes].items.catch(monkey.items.throw())
				} else {
					monkeys[monkey.test.no].items.catch(monkey.items.throw())
				}
			}
		}
	}
	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Ints(inspections)
	monkeyBusiness := inspections[len(inspections)-1] * inspections[len(inspections)-2]
	return monkeyBusiness
}

type Monkeys []*Monkey

type Monkey struct {
	inspections int
	items       Items
	op          []int
	test        Test
}

type Test struct {
	value int
	yes   int
	no    int
}

type Items []*int

func (c *Items) catch(i *int) { *c = append(*c, i) }
func (c *Items) throw() *int  { i := (*c)[0]; *c = (*c)[1:]; return i }

func (m *Monkey) inspect(item *int) {
	m.inspections++
	switch m.op[0] {
	case 0:
		*item += m.op[1]
	case 1:
		*item *= m.op[1]
	case 2:
		*item = int(math.Pow(float64(*item), float64(m.op[1])))
	}
}

func lowestCommonMultiple(a, b int) int {
	lcm := a * b / greatestCommonDenominator(a, b)
	return lcm
}

func greatestCommonDenominator(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func parseMonkeys(input *[]byte) Monkeys {
	monkeys := Monkeys{}
	for _, block := range strings.Split(strings.TrimSpace(string(*input)), "\n\n") {
		monkey := &Monkey{}
		monkey.inspections = 0
		for i, line := range strings.Split(strings.TrimSpace(block), "\n") {
			if i == 0 {
				continue
			}
			properties := strings.Split(strings.TrimSpace(line), ":")
			switch properties[0] {
			case "Starting items":
				items := strings.Split(properties[1], ",")
				for _, item := range items {
					worryLevel, err := strconv.Atoi(strings.TrimSpace(item))
					if err != nil {
						logrus.Error(err)
					}
					item := int(worryLevel)
					monkey.items.catch(&item)
				}
			case "Operation":
				operation := strings.Split(strings.TrimSpace(properties[1]), " ")
				switch operation[3] {
				case "+":
					monkey.op = append(monkey.op, 0)
				case "*":
					monkey.op = append(monkey.op, 1)
				}
				if operation[4] == "old" {
					monkey.op[0] = 2
					monkey.op = append(monkey.op, 2)
				} else {
					op, err := strconv.Atoi(operation[4])
					if err != nil {
						logrus.Error(err)
					}
					operation := int(op)
					monkey.op = append(monkey.op, operation)
				}
			case "Test":
				test := strings.Split(strings.TrimSpace(properties[1]), " ")
				value, err := strconv.Atoi(test[2])
				if err != nil {
					logrus.Error(err)
				}
				monkey.test.value = int(value)

			case "If true":
				throwTo := strings.Split(strings.TrimSpace(properties[1]), " ")
				newMonkey, err := strconv.Atoi(throwTo[3])
				if err != nil {
					logrus.Error(err)
				}
				monkey.test.yes = newMonkey
			case "If false":
				throwTo := strings.Split(strings.TrimSpace(properties[1]), " ")
				newMonkey, err := strconv.Atoi(throwTo[3])
				if err != nil {
					logrus.Error(err)
				}
				monkey.test.no = newMonkey
			}
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}
