package day1

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
R102`

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 3
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 7
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
