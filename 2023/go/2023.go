package y23

import (
	"fmt"
	"io/ioutil"
	"github.com/jrhorner1/AoC/2023/go/day1"
	"github.com/jrhorner1/AoC/2023/go/day2"
	log "github.com/sirupsen/logrus"
)

func Run(year, day *int) {
	switch *day {
	case 1:
		input, _ := ioutil.ReadFile("2023/input/1")
		output(year, day, 1, fmt.Sprint(day1.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day1.Puzzle(&input, true)))
	case 2:
		input, _ := ioutil.ReadFile("2023/input/2")
		output(year, day, 1, fmt.Sprint(day2.Puzzle(&input, false)))
		output(year, day, 2, fmt.Sprint(day2.Puzzle(&input, true)))
	default:
		panic("unimplemented")
	case 26, 27, 28, 29, 30, 31:
		panic("What did you expect to find here? It's an advent calendar.")
	}
}

func output(year, day *int, part int, solution string) {
	log.Infof("%d - Day %d - Part %d: %s", *year, *day, part, solution)
}
