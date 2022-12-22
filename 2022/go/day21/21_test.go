package day21

import (
	"testing"

	"github.com/sirupsen/logrus"
)

***REMOVED***

func Test_p1(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(152)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(301)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
