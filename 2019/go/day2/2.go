package day2

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

var filename = "2019/input/2"

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
	// Restore gravity assist program to "1202 program alarm" state
	program[1], program[2] = 12, 2
	computer := intcode.NewComputer(&program)
	computer.Run()
	output := computer.GetMemory()

	if part2 {
		searchValue := 19690720
		var noun, verb int
		for noun = 0; noun < 100; noun++ {
			for verb = 0; verb < 100; verb++ {
				program[1], program[2] = noun, verb
				computer = intcode.NewComputer(&program)
				computer.Run()
				output = computer.GetMemory()
				if (*output)[0] == searchValue {
					break
				}
			}
			if (*output)[0] == searchValue {
				break
			}
		}
		return 100*noun + verb
	}
	return (*output)[0]
}
