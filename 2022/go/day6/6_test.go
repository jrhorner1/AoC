package day6

import (
	"testing"
)

var (
	examples = []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb\n",
		"bvwbjplbgvbhsrlpgdmjqwftvncz\n",
		"nppdvjthqldpwncqszvftbrmjlhg\n",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg\n",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw\n",
	}
	answers = []int{
		7,
		5,
		6,
		10,
		11,
		// part 2 answers
		19,
		23,
		23,
		29,
		26,
	}
	index = 5
)

func Test_p1(t *testing.T) {
	for i, example := range examples {
		input := []byte(example)
		got := Puzzle(&input, false)
		want := answers[i]
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
}

func Test_p2(t *testing.T) {
	for i, example := range examples {
		input := []byte(example)
		got := Puzzle(&input, true)
		want := answers[i+index]
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
}
