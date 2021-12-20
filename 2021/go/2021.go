package y21

import (
	"fmt"
	"io/ioutil"

	"github.com/jrhorner1/AoC/2021/go/day1"
	"github.com/jrhorner1/AoC/2021/go/day10"
	"github.com/jrhorner1/AoC/2021/go/day11"
	"github.com/jrhorner1/AoC/2021/go/day12"
	"github.com/jrhorner1/AoC/2021/go/day13"
	"github.com/jrhorner1/AoC/2021/go/day14"
	"github.com/jrhorner1/AoC/2021/go/day15"
	"github.com/jrhorner1/AoC/2021/go/day16"
	"github.com/jrhorner1/AoC/2021/go/day17"
	"github.com/jrhorner1/AoC/2021/go/day2"
	"github.com/jrhorner1/AoC/2021/go/day3"
	"github.com/jrhorner1/AoC/2021/go/day4"
	"github.com/jrhorner1/AoC/2021/go/day5"
	"github.com/jrhorner1/AoC/2021/go/day6"
	"github.com/jrhorner1/AoC/2021/go/day7"
	"github.com/jrhorner1/AoC/2021/go/day8"
	"github.com/jrhorner1/AoC/2021/go/day9"
)

func Run(year, day *int) {
	switch *day {
	case 1:
		input, _ := ioutil.ReadFile("2021/input/1")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day1.Puzzle(&input, false), day1.Puzzle(&input, true))
	case 2:
		input, _ := ioutil.ReadFile("2021/input/2")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day2.Puzzle(&input, false), day2.Puzzle(&input, true))
	case 3:
		input, _ := ioutil.ReadFile("2021/input/3")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day3.Puzzle(&input, false), day3.Puzzle(&input, true))
	case 4:
		input, _ := ioutil.ReadFile("2021/input/4")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day4.Puzzle(&input, false), day4.Puzzle(&input, true))
	case 5:
		input, _ := ioutil.ReadFile("2021/input/5")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day5.Puzzle(&input, false), day5.Puzzle(&input, true))
	case 6:
		input, _ := ioutil.ReadFile("2021/input/6")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day6.Puzzle(&input, 80), day6.Puzzle(&input, 256))
	case 7:
		input, _ := ioutil.ReadFile("2021/input/7")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day7.Puzzle(&input, false), day7.Puzzle(&input, true))
	case 8:
		input, _ := ioutil.ReadFile("2021/input/8")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day8.Puzzle(&input, false), day8.Puzzle(&input, true))
	case 9:
		input, _ := ioutil.ReadFile("2021/input/9")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day9.Puzzle(&input, false), day9.Puzzle(&input, true))
	case 10:
		input, _ := ioutil.ReadFile("2021/input/10")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day10.Puzzle(&input, false), day10.Puzzle(&input, true))
	case 11:
		input, _ := ioutil.ReadFile("2021/input/11")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day11.Puzzle(&input, 100), day11.Puzzle(&input, 10000))
	case 12:
		input, _ := ioutil.ReadFile("2021/input/12")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day12.Puzzle(&input, false), day12.Puzzle(&input, true))
	case 13:
		input, _ := ioutil.ReadFile("2021/input/13")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: \n", *year, *day, day13.Puzzle(&input, false))
		day13.Puzzle(&input, true)
	case 14:
		input, _ := ioutil.ReadFile("2021/input/14")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day14.Puzzle(&input, 10), day14.Puzzle(&input, 40))
	case 15:
		input, _ := ioutil.ReadFile("2021/input/15")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day15.Puzzle(&input, false), day15.Puzzle(&input, true))
	case 16:
		input, _ := ioutil.ReadFile("2021/input/16")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day16.Puzzle(&input, false), day16.Puzzle(&input, true))
	case 17:
		input, _ := ioutil.ReadFile("2021/input/17")
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: %d\n", *year, *day, day17.Puzzle(&input, false), day17.Puzzle(&input, true))
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}
