package y22

import (
	"fmt"
	"io/ioutil"

	"github.com/jrhorner1/AoC/2022/go/day1"
	"github.com/jrhorner1/AoC/2022/go/day10"
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
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day1.Puzzle(&input, false))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day1.Puzzle(&input, true))
	case 2:
		input, _ := ioutil.ReadFile("2022/input/2")
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day2.Puzzle(&input, false))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day2.Puzzle(&input, true))
	case 3:
		input, _ := ioutil.ReadFile("2022/input/3")
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day3.Puzzle(&input, false))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day3.Puzzle(&input, true))
	case 4:
		input, _ := ioutil.ReadFile("2022/input/4")
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day4.Puzzle(&input, false))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day4.Puzzle(&input, true))
	case 5:
		input, _ := ioutil.ReadFile("2022/input/5")
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day5.Puzzle(&input, false))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day5.Puzzle(&input, true))
	case 6:
		input, _ := ioutil.ReadFile("2022/input/6")
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day6.Puzzle(&input, 4))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day6.Puzzle(&input, 14))
	case 7:
		input, _ := ioutil.ReadFile("2022/input/7")
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day7.Puzzle(&input, false))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day7.Puzzle(&input, true))
	case 8:
		input, _ := ioutil.ReadFile("2022/input/8")
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day8.Puzzle(&input, false))
		log.Infof("%d - Day %d - Part 1: %v", *year, *day, day8.Puzzle(&input, true))
	case 9:
		input, _ := ioutil.ReadFile("2022/input/9")
		output(year, day, 1, fmt.Sprint(day9.Puzzle(&input, 2)))
		output(year, day, 2, fmt.Sprint(day9.Puzzle(&input, 10)))
	case 10:
		input, _ := ioutil.ReadFile("2022/input/10")
		p1, p2 := day10.Puzzle(&input)
		fmt.Printf("\t%d Day %d solutions\nPart 1: %d\nPart 2: \n\n%v\n", *year, *day, p1, p2)
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}

func output(year, day *int, part int, solution string) {
	log.Infof("%d - Day %d - Part %d: %s", *year, *day, part, solution)
}
