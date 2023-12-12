package day10

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const example1 string = `.....
.S-7.
.|.|.
.L-J.
.....
`

func Test_p1_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example1)
	got := Puzzle(&input, false)
	want := 4
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

const example2 string = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

func Test_p1_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example2)
	got := Puzzle(&input, false)
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

const example3 string = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`

func Test_p2_1(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example3)
	got := Puzzle(&input, true)
	want := 4
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

const example4 string = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

func Test_p2_2(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example4)
	got := Puzzle(&input, true)
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

const example5 string = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`

func Test_p2_3(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	input := []byte(example5)
	got := Puzzle(&input, true)
	want := 10
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
