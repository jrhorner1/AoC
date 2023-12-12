package day10

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

***REMOVED***

func Test_p1_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example1)
	got := Puzzle(&input, false)
	want := 4
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

***REMOVED***

func Test_p1_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, false)
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

***REMOVED***

func Test_p2_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example3)
	got := Puzzle(&input, true)
	want := 4
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

***REMOVED***

func Test_p2_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example4)
	got := Puzzle(&input, true)
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

***REMOVED***

func Test_p2_3(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example5)
	got := Puzzle(&input, true)
	want := 10
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
