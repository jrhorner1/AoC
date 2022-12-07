package day7

import (
	"testing"
)

const example string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func Test_p1(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, false)
	want := int(95437)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test_p2(t *testing.T) {
	input := []byte(example)
	got := Puzzle(&input, true)
	want := int(24933642)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
