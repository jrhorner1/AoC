package day5

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := string("CMZ")
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := string("MCD")
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
