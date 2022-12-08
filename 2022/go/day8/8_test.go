package day8

import (
	"testing"
)

const example string = `30373
25512
65332
33549
35390
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(21)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(8)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
