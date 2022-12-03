package day1

import (
	"testing"
)

const example string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(24000)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(45000)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
