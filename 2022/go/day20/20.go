package day20

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var decryptionKey = 811589153

func Puzzle(input *[]byte, count int) int {
	order, zero := parseDigits(input, count)
	mixed := Digits{}
	for _, digit := range order {
		mixed = append(mixed, digit)
	}
	for loop := 0; loop < count; loop++ {
		for _, next := range order {
			logrus.Debugf("Next: %d", *next)
			for i, digit := range mixed {
				if digit == next {
					idx := (i + *digit) % (len(mixed) - 1)
					for idx < 0 {
						idx += len(mixed) - 1
					}
					if idx == 0 {
						idx = len(mixed) - 1
					}
					logrus.Debugf("Digit: %d Index %d New Index: %d", *digit, i, idx)
					if idx > i {
						for a, b := i, i+1; a < idx; a, b = a+1, b+1 {
							mixed[a], mixed[b] = mixed[b], mixed[a]
						}
						i -= 1
					} else if idx < i {
						for a, b := i-1, i; a >= idx; a, b = a-1, b-1 {
							mixed[a], mixed[b] = mixed[b], mixed[a]
						}
					}
					break
				}
			}
		}
	}
	list := []int{}
	for i := range mixed {
		list = append(list, *(mixed[i]))
	}
	logrus.Debug(list)
	list = []int{}
	for i := range order {
		list = append(list, *(order[i]))
	}
	logrus.Debug(list)
	zeroidx := 0
	for i, digit := range mixed {
		if digit == zero {
			zeroidx = i
			break
		}
	}
	logrus.Debugf("Zero Index: %d", zeroidx)
	xidx, yidx, zidx := zeroidx+(1000%len(mixed)), zeroidx+(2000%len(mixed)), zeroidx+(3000%len(mixed))
	for _, idx := range []*int{&xidx, &yidx, &zidx} {
		if *idx > len(mixed)-1 {
			*idx -= len(mixed)
		}
	}
	logrus.Debugf("xidx: %d yidx: %d zidx: %d", xidx, yidx, zidx)
	x, y, z := mixed[xidx], mixed[yidx], mixed[zidx]
	logrus.Infof("x: %d y: %d z: %d", *x, *y, *z)
	sum := *x + *y + *z
	logrus.Debugf("Grove coordinates sum: %d", sum)
	return sum
}

type Digits []*int

func parseDigits(input *[]byte, loops int) (Digits, *int) {
	encrypted := Digits{}
	var zero *int
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		value, err := strconv.Atoi(line)
		if err != nil {
			logrus.Error(err)
		}
		if loops > 1 {
			value *= decryptionKey
		}
		encrypted = append(encrypted, &value)
		if value == 0 {
			zero = encrypted[len(encrypted)-1]
		}
	}
	return encrypted, zero
}
