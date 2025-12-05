package day3

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example = `987654321111111
811111111111119
234234234234278
818181911112111`

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 357
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 3121910778619
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
