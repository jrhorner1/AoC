package day25

import (
	"testing"
)

const example string = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 58
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
