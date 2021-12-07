package day2

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
