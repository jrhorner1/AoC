package day1

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

***REMOVED***

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example1)
	got := Puzzle(&input, false)
	want := 142
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

***REMOVED***

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, true)
	want := 281
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
