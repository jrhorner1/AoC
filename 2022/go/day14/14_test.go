package day14

import (
	"testing"
)

const example string = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(24)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(93)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
