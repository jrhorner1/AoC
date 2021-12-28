package day22

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example1)
	got := Puzzle(&input, true)
	want := 590784
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

***REMOVED***

func Test_p2(t *testing.T) {
	input := []byte(example2)
	got := Puzzle(&input, false)
	want := 2758514936282235
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
