package day21

import (
	"testing"
)

const example string = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(152)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(301)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
