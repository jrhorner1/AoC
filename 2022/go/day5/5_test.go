package day5

import (
	"testing"
)

const example string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

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
