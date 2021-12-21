package day21

import (
	"fmt"
	"testing"
)

const example string = "Player 1 starting position: 4\nPlayer 2 starting position: 8\n"

func Test_deterministic(t *testing.T) {
	input := []byte(fmt.Sprintf(example))
	p1, p2 := parseInput(&input)
	rolls := 0
	deterministicDirac(&p1, &p2, 100, &rolls)
	got := p2.score * rolls
	want := 739785
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_quantum(t *testing.T) {
	input := []byte(fmt.Sprintf(example))
	p1, p2 := parseInput(&input)
	p1.wins, p2.wins = quantumDirac(p1, p2)
	got := p1.wins
	want := 444356092776315
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
