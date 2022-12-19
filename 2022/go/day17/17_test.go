package day17

import (
	"testing"
)

const example string = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, 2022)
	want := int(3068)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, 1000000000000)
	want := int(1514285714288)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

/*
8170 days to calculate by iterating through each step using bitfield
*/
