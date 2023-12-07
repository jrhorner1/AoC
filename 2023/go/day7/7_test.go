package day7

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func Test_p1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, false)
	want := 6440
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example)
	got := Puzzle(&input, true)
	want := 5905
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

// https://www.reddit.com/r/adventofcode/comments/18cr4xr/2023_day_7_better_example_input_not_a_spoiler/
const example2 string = `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41
`

func Test_p2_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, true)
	want := 6839
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
