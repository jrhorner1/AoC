package day10

import (
	"testing"
)

***REMOVED***

const answer2 = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got, _ := Puzzle(&input)
	want := int(13140)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	_, got := Puzzle(&input)
	want := answer2
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
