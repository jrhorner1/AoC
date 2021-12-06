package main

import (
	"flag"
	"time"

	y21 "github.com/jrhorner1/AoC/2021/go"
)

func main() {
	var year, day int
	flag.IntVar(&year, "y", time.Now().Year(), "Year to run a puzzle solution from. Defaults to the current year.")
	flag.IntVar(&day, "d", time.Now().Day(), "Day to run a puzzle solution from. Defaults to the current day.")
	flag.Parse()

	switch year {
	case 2021:
		run2021(&day)
	}
}

func run2021(day *int) {
	switch *day {
	case 1:
		y21.Day1()
	}
}
