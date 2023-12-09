package day8

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example1 string = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

const example2 string = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

func Test_p1_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example1)
	got := Puzzle(&input, false)
	want := int64(2)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func Test_p1_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, false)
	want := int64(6)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

const example3 string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func Test_p2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example3)
	got := Puzzle(&input, true)
	want := int64(6)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
