package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func main() {
	input, _ := ioutil.ReadFile("2020/input/8")
	instr := strings.Split(strings.TrimSpace(string(input)), "\n")
	acc, _ := run(instr)
	fmt.Println("Part 1:", acc)
	for i, s := range instr {
		n_instr := make([]string, len(instr))
		copy(n_instr, instr)
		n_instr[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(s)
		if acc, inf := run(n_instr); inf == false {
			fmt.Println("Part 2:", acc)
		}
	}
}
