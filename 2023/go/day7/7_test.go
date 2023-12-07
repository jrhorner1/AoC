package day7

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

***REMOVED***

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 6440
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 5905
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

// https://www.reddit.com/r/adventofcode/comments/18cr4xr/2023_day_7_better_example_input_not_a_spoiler/
***REMOVED***

func Test_p2_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, true)
	want := 6839
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
