package day9

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

var filename = "2019/input/9"

func Part1() int {
	return puzzle(Input(filename), 1)
}

func Part2() int {
	return puzzle(Input(filename), 2)
}

func Input(file string) []int {
	input, _ := ioutil.ReadFile(file)
	in := strings.Split(strings.TrimSpace(string(input)), ",")
	var output []int
	for _, i := range in {
		out, _ := strconv.Atoi(i)
		output = append(output, out)
	}
	return output
}

func puzzle(input []int, init int) int {
	program := Input(filename)
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
