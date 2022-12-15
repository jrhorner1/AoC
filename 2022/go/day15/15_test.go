package day15

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false, 10)
	want := int(26)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true, 10)
	want := int(56000011)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
