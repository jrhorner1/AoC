package day20

import (
	"testing"

	"github.com/sirupsen/logrus"
)

const example string = `1
2
-3
3
-2
0
4
`

func Test_p1(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, 1)
	want := int(3)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, 10)
	want := int(1623178306)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
