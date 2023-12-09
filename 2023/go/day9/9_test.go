package day9

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 114
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
