package day1

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example1 string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example1)
	got := Puzzle(&input, false)
	want := 142
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

const example2 string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, true)
	want := 281
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
