package y16

import (
	"fmt"
	"io/ioutil"

	"github.com/jrhorner1/AoC/2016/go/day1"
	"github.com/jrhorner1/AoC/2016/go/day2"
	"github.com/jrhorner1/AoC/2016/go/day3"
)

func Run(year, day *int) {
	switch *day {
	case 1:
		input, _ := ioutil.ReadFile("2016/input/1")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day1.Puzzle(&input, false), day1.Puzzle(&input, true))
	case 2:
		input, _ := ioutil.ReadFile("2016/input/2")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %s\nPart 2: %s\n", *year, *day, day2.Puzzle(&input, false), day2.Puzzle(&input, true))
	case 3:
		input, _ := ioutil.ReadFile("2016/input/3")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day3.Puzzle(&input, false), day3.Puzzle(&input, true))
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}
