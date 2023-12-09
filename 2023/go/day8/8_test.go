package day8

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

***REMOVED***

***REMOVED***

func Test_p1_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example1)
	got := Puzzle(&input, false)
	want := int64(2)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func Test_p1_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, false)
	want := int64(6)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

***REMOVED***

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example3)
	got := Puzzle(&input, true)
	want := int64(6)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
