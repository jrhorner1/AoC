package day22

import (
	"testing"
)

const example string = `        ...#
.#..
#...
....
...#.......#
........#...
..#....#....
..........#.
...#....
.....#..
.#......
......#.

10R5L5R10L4R5L5
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(6032)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 42
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
