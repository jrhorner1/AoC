package day23

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 12521
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

const insert string = `  #D#C#B#A#
  #D#B#A#C#`

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 44169
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
