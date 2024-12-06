package day5

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example = ``

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 143
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 123
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
