package day9

import (
	"testing"
)

const (
	example1 string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
	example2 string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
)

func Test_p1(t *testing.T) {
	input := []byte(example1)
	got := Puzzle(&input, 2)
	want := int(13)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example2)
	got := Puzzle(&input, 10)
	want := int(36)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
