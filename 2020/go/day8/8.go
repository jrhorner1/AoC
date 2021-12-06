package day8

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2020/input/8"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func run(instr []string) (int, bool) {
	acc, steps := 0, 0
	seen := map[int]bool{}
	for steps < len(instr) {
		if seen[steps] {
			return acc, true
		} else {
			seen[steps] = true
		}
		cmd := strings.Split(instr[steps], " ")[0]
		val, _ := strconv.Atoi(strings.Split(instr[steps], " ")[1])
		switch cmd {
		case "acc":
			acc += val
		case "jmp":
			steps += val - 1
		}
		steps++
	}
	return acc, false
}

func puzzle(input []string, part2 bool) int {
	if part2 {
		for i, s := range input {
			n_instr := make([]string, len(input))
			copy(n_instr, input)
			n_instr[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(s)
			if acc, inf := run(n_instr); inf == false {
				return acc
			}
		}
	}
	acc, _ := run(input)
	return acc
}
