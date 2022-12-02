package y22

import (
	"fmt"
	"io/ioutil"
	"github.com/jrhorner1/AoC/2022/go/day1"
	"github.com/jrhorner1/AoC/2022/go/day2"
)

func Run(year, day *int) {
	switch *day {
	case 1:
		input, _ := ioutil.ReadFile("2022/input/1")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day1.Puzzle(&input, false), day1.Puzzle(&input, true))
	case 2:
		input, _ := ioutil.ReadFile("2022/input/2")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day2.Puzzle(&input, false), day2.Puzzle(&input, true))
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}
