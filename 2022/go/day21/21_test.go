package day21

import (
	"testing"

	"github.com/sirupsen/logrus"
)

***REMOVED***

/*
            root
           /   \
         pppw sjmn
        /   | |   \
    cczh lfql drzm dbpl
   /   |      |   \
sllz lgvd     hmdt sczc
    /   |
ljgn ptdq
    /   |
humn dvpt
*/

func Test_p1(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := uint64(152)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := uint64(301)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
