package day9

import (
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

func Puzzle(input *[]byte, init int) int {
	var program []int
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), ",") {
		out, _ := strconv.Atoi(i)
		program = append(program, out)
	}
	computer := intcode.NewComputer(&program)
	go computer.Run()
	computer.Input <- init
	channelOpen := true
	var output int
	for channelOpen {
		output, channelOpen = <-computer.Output
		if channelOpen {
			break
		}
	}
	return output
}
