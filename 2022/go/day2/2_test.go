package day2

import (
	"testing"
)

const example string = `A Y
B X
C Z
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(15)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(12)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
