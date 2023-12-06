package day6

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example string = `Time:      7  15   30
Distance:  9  40  200
`

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 288
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 71503
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
