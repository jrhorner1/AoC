package day25

import (
	"testing"
)

***REMOVED***

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 58
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
