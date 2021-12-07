package day8

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n")
	if part2 {
		for i, s := range in {
			n_instr := make([]string, len(in))
			copy(n_instr, in)
			n_instr[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(s)
			if acc, inf := run(n_instr); inf == false {
				return acc
			}
		}
	}
	acc, _ := run(in)
	return acc
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
