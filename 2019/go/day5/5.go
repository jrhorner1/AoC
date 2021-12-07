package day5

import (
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

func Puzzle(input *[]byte, part2 bool) int {
	var program []int
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), ",") {
		tmp, _ := strconv.Atoi(i)
		program = append(program, tmp)
	}
	computer := intcode.NewComputer(&program)
	if part2 {
		computer.Input <- 5
	} else {
		computer.Input <- 1
	}
	go computer.Run()
	var output int
	for {
		if out, openChannel := <-computer.Output; openChannel {
			if out > output {
				output = out
			}
		} else {
			break
		}
	}
	return output
}
