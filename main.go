/*?sr/bin/env go run "$0" "$@"; exit $? #*/
// This is actually not a shebang, the first line is both valid shell script and valid go code.

package main

import (
	"flag"
	"time"

	y15 "github.com/jrhorner1/AoC/2015/go"
	y16 "github.com/jrhorner1/AoC/2016/go"
	y17 "github.com/jrhorner1/AoC/2017/go"
	y18 "github.com/jrhorner1/AoC/2018/go"
	y19 "github.com/jrhorner1/AoC/2019/go"
	y20 "github.com/jrhorner1/AoC/2020/go"
	y21 "github.com/jrhorner1/AoC/2021/go"
	y22 "github.com/jrhorner1/AoC/2022/go"
	"github.com/jrhorner1/AoC/pkg/math"
	log "github.com/sirupsen/logrus"
)

func main() {
	start := time.Now()
	var year, day, executions int
	flag.IntVar(&year, "y", time.Now().Year(), "Year to run a puzzle solution from. Defaults to the current year.")
	flag.IntVar(&day, "d", time.Now().Day(), "Day to run a puzzle solution from. Defaults to the current day.")
	flag.IntVar(&executions, "a", 1, "Average execution time to solve a puzzle N times. Default is 1.")
	flag.Parse()
	averages := []int64{}
	for i := 0; i < executions; i++ {
		execStart := time.Now()
		switch year {
		case 2015:
			y15.Run(&year, &day)
		case 2016:
			y16.Run(&year, &day)
		case 2017:
			y17.Run(&year, &day)
		case 2018:
			y18.Run(&year, &day)
		case 2019:
			y19.Run(&year, &day)
		case 2020:
			y20.Run(&year, &day)
		case 2021:
			y21.Run(&year, &day)
		case 2022:
			y22.Run(&year, &day)
		default:
			panic("Either advent of code didn't exist, or the year you're looking for isn't here yet.")
		}
		averages = append(averages, time.Since(execStart).Nanoseconds())
	}
	if executions > 1 {
		average := time.Duration(math.Average(&averages))
		log.Infof("Average Execution time out of %d: %v", executions, average.String())
	}
	log.Infof("Execution time: %v", time.Since(start).String())
}
