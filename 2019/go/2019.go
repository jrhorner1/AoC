package y19

import (
	"fmt"
	"io/ioutil"

	"github.com/jrhorner1/AoC/2019/go/day1"
	"github.com/jrhorner1/AoC/2019/go/day10"
	"github.com/jrhorner1/AoC/2019/go/day11"
	"github.com/jrhorner1/AoC/2019/go/day12"
	"github.com/jrhorner1/AoC/2019/go/day13"
	"github.com/jrhorner1/AoC/2019/go/day2"
	"github.com/jrhorner1/AoC/2019/go/day3"
	"github.com/jrhorner1/AoC/2019/go/day4"
	"github.com/jrhorner1/AoC/2019/go/day5"
	"github.com/jrhorner1/AoC/2019/go/day6"
	"github.com/jrhorner1/AoC/2019/go/day7"
	"github.com/jrhorner1/AoC/2019/go/day8"
	"github.com/jrhorner1/AoC/2019/go/day9"
)

func Run(year, day *int) {
	switch *day {
	case 1:
		input, _ := ioutil.ReadFile("2019/input/1")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day1.Puzzle(&input, false), day1.Puzzle(&input, true))
	case 2:
		input, _ := ioutil.ReadFile("2019/input/2")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day2.Puzzle(&input, false), day2.Puzzle(&input, true))
	case 3:
		input, _ := ioutil.ReadFile("2019/input/3")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day3.Puzzle(&input, false), day3.Puzzle(&input, true))
	case 4:
		input, _ := ioutil.ReadFile("2019/input/4")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day4.Puzzle(&input, false), day4.Puzzle(&input, true))
	case 5:
		input, _ := ioutil.ReadFile("2019/input/5")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day5.Puzzle(&input, false), day5.Puzzle(&input, true))
	case 6:
		input, _ := ioutil.ReadFile("2019/input/6")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day6.Puzzle(&input, false), day6.Puzzle(&input, true))
	case 7:
		input, _ := ioutil.ReadFile("2019/input/7")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day7.Puzzle(&input, []int{0, 1, 2, 3, 4}), day7.Puzzle(&input, []int{5, 6, 7, 8, 9}))
	case 8:
		input, _ := ioutil.ReadFile("2019/input/8")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2:\n", *year, *day, day8.Puzzle(&input, false))
		day8.Puzzle(&input, true)
	case 9:
		input, _ := ioutil.ReadFile("2019/input/9")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day9.Puzzle(&input, 1), day9.Puzzle(&input, 2))
	case 10:
		input, _ := ioutil.ReadFile("2019/input/10")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day10.Puzzle(&input, false), day10.Puzzle(&input, true))
	case 11:
		input, _ := ioutil.ReadFile("2019/input/11")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2:\n", *year, *day, day11.Puzzle(&input, 0))
		day11.Puzzle(&input, 1)
	case 12:
		input, _ := ioutil.ReadFile("2019/input/12")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day12.Puzzle(&input, false), day12.Puzzle(&input, true))
	case 13:
		input, _ := ioutil.ReadFile("2019/input/13")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day13.Puzzle(&input, false), day13.Puzzle(&input, true))
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}
