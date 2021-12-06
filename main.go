package main

import (
	"flag"
	"fmt"
	"time"

	y15 "github.com/jrhorner1/AoC/2015/go"
	y16 "github.com/jrhorner1/AoC/2016/go"
	y17 "github.com/jrhorner1/AoC/2017/go"
	y18 "github.com/jrhorner1/AoC/2018/go"
	y19 "github.com/jrhorner1/AoC/2019/go"
	y20 "github.com/jrhorner1/AoC/2020/go"
	y21 "github.com/jrhorner1/AoC/2021/go"
)

func main() {
	var year, day int
	flag.IntVar(&year, "y", time.Now().Year(), "Year to run a puzzle solution from. Defaults to the current year.")
	flag.IntVar(&day, "d", time.Now().Day(), "Day to run a puzzle solution from. Defaults to the current day.")
	flag.Parse()

	switch year {
	case 2015:
		y15.Run(&day)
	case 2016:
		y16.Run(&day)
	case 2017:
		y17.Run(&day)
	case 2018:
		y18.Run(&day)
	case 2019:
		y19.Run(&day)
	case 2020:
		y20.Run(&day)
	case 2021:
		y21.Run(&day)
	default:
		fmt.Println("The year either hasn't been implmented yet or advent of code didnt exist that year.")
	}

}
