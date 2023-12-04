package day3

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 4361
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 467835
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
