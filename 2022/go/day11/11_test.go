package day11

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(10605)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(2713310158)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
