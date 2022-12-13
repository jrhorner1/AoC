package y22

import (
	"fmt"
	"io/ioutil"

	"github.com/jrhorner1/AoC/2022/go/day1"
	"github.com/jrhorner1/AoC/2022/go/day10"
	"github.com/jrhorner1/AoC/2022/go/day11"
	"github.com/jrhorner1/AoC/2022/go/day12"
	"github.com/jrhorner1/AoC/2022/go/day13"
	"github.com/jrhorner1/AoC/2022/go/day2"
	"github.com/jrhorner1/AoC/2022/go/day3"
	"github.com/jrhorner1/AoC/2022/go/day4"
	"github.com/jrhorner1/AoC/2022/go/day5"
	"github.com/jrhorner1/AoC/2022/go/day6"
	"github.com/jrhorner1/AoC/2022/go/day7"
	"github.com/jrhorner1/AoC/2022/go/day8"
	"github.com/jrhorner1/AoC/2022/go/day9"
	log "github.com/sirupsen/logrus"
)

func Run(year, day *int) {
	switch *day {
	case 1:
		input, _ := ioutil.ReadFile("2022/input/1")
		output(year, day, 1, fmt.Sprint(day1.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day1.Puzzle(&input, true)))
	case 2:
		input, _ := ioutil.ReadFile("2022/input/2")
		output(year, day, 1, fmt.Sprint(day2.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day2.Puzzle(&input, true)))
	case 3:
		input, _ := ioutil.ReadFile("2022/input/3")
		output(year, day, 1, fmt.Sprint(day3.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day4.Puzzle(&input, true)))
	case 4:
		input, _ := ioutil.ReadFile("2022/input/4")
		output(year, day, 1, fmt.Sprint(day4.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day4.Puzzle(&input, true)))
	case 5:
		input, _ := ioutil.ReadFile("2022/input/5")
		output(year, day, 1, fmt.Sprint(day5.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day5.Puzzle(&input, true)))
	case 6:
		input, _ := ioutil.ReadFile("2022/input/6")
		output(year, day, 1, fmt.Sprint(day6.Puzzle(&input, 4)))
		output(year, day, 2, fmt.Sprint(day6.Puzzle(&input, 14)))
	case 7:
		input, _ := ioutil.ReadFile("2022/input/7")
		output(year, day, 1, fmt.Sprint(day7.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day7.Puzzle(&input, true)))
	case 8:
		input, _ := ioutil.ReadFile("2022/input/8")
		output(year, day, 1, fmt.Sprint(day8.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day8.Puzzle(&input, true)))
	case 9:
		input, _ := ioutil.ReadFile("2022/input/9")
		output(year, day, 1, fmt.Sprint(day9.Puzzle(&input, 2)))
		output(year, day, 2, fmt.Sprint(day9.Puzzle(&input, 10)))
	case 10:
		input, _ := ioutil.ReadFile("2022/input/10")
		p1, p2 := day10.Puzzle(&input)
		output(year, day, 1, fmt.Sprint(p1))
		output(year, day, 2, fmt.Sprint("\n", p2))
	case 11:
		input, _ := ioutil.ReadFile("2022/input/11")
		output(year, day, 1, fmt.Sprint(day11.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day11.Puzzle(&input, true)))
	case 12:
		input, _ := ioutil.ReadFile("2022/input/12")
		output(year, day, 1, fmt.Sprint(day12.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day12.Puzzle(&input, true)))
	case 13:
		input, _ := ioutil.ReadFile("2022/input/13")
		output(year, day, 1, fmt.Sprint(day13.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day13.Puzzle(&input, true)))
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}

func output(year, day *int, part int, solution string) {
	log.Infof("%d Day %d|Part %d: %s", *year, *day, part, solution)
}
