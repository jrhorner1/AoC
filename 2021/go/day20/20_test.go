package day20

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, 2)
	want := 35
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, 50)
	want := 3351
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
