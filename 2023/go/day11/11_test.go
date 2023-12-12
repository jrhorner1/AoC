package day11

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

***REMOVED***

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, 2)
	want := 374
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, 10)
	want := 1030
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, 100)
	want := 8410
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
