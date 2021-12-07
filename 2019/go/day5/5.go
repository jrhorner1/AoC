package day5

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

var filename = "2019/input/5"

// var filename = "2019/examples/5"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), ",")
	return output
}

func puzzle(input []string, part2 bool) int {
	var program []int
	for i := 0; i < len(input); i++ {
		tmp, _ := strconv.Atoi(input[i])
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
