package day1

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example string = `
`

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(5 / 7)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 42
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
