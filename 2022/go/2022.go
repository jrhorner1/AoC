package y22

import (
	"fmt"
	"io/ioutil"

	"github.com/jrhorner1/AoC/2022/go/day1"
	"github.com/jrhorner1/AoC/2022/go/day2"
	"github.com/jrhorner1/AoC/2022/go/day3"
	"github.com/jrhorner1/AoC/2022/go/day4"
	"github.com/jrhorner1/AoC/2022/go/day5"
	"github.com/jrhorner1/AoC/2022/go/day6"
	"github.com/jrhorner1/AoC/2022/go/day7"
)

func Run(year, day *int) {
	switch *day {
	case 1:
		input, _ := ioutil.ReadFile("2022/input/1")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day1.Puzzle(&input, false), day1.Puzzle(&input, true))
	case 2:
		input, _ := ioutil.ReadFile("2022/input/2")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day2.Puzzle(&input, false), day2.Puzzle(&input, true))
	case 3:
		input, _ := ioutil.ReadFile("2022/input/3")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day3.Puzzle(&input, false), day3.Puzzle(&input, true))
	case 4:
		input, _ := ioutil.ReadFile("2022/input/4")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day4.Puzzle(&input, false), day4.Puzzle(&input, true))
	case 5:
		input, _ := ioutil.ReadFile("2022/input/5")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %s\nPart 2: %s\n", *year, *day, day5.Puzzle(&input, false), day5.Puzzle(&input, true))
	case 6:
		input, _ := ioutil.ReadFile("2022/input/6")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day6.Puzzle(&input, 4), day6.Puzzle(&input, 14))
	case 7:
		input, _ := ioutil.ReadFile("2022/input/7")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day7.Puzzle(&input, false), day7.Puzzle(&input, true))
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}
