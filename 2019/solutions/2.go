package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

func main() {
	input, _ := ioutil.ReadFile("2019/input/2")
	in := strings.Split(strings.TrimSpace(string(input)), ",")
	var program []int
	for i := 0; i < len(in); i++ {
		tmp, _ := strconv.Atoi(in[i])
		program = append(program, tmp)
	}

	// Restore gravity assist program to "1202 program alarm" state
	program[1], program[2] = 12, 2
	ic := intcode.NewComputer(&program)
	ic.Run()
	output := ic.GetMemory()
	fmt.Println("Part 1:", (*output)[0])

	searchValue := 19690720
	var noun, verb int
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			program[1], program[2] = noun, verb
			ic = intcode.NewComputer(&program)
			ic.Run()
			output = ic.GetMemory()
			if (*output)[0] == searchValue {
				break
			}
		}
		if (*output)[0] == searchValue {
			break
		}
	}
	fmt.Println("Part 2:", 100*noun+verb)
}
