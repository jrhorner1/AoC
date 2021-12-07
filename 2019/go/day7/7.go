package day7

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

var filename = "2019/input/7"

func Part1() int {
	phaseVals := []int{0, 1, 2, 3, 4}
	return puzzle(input(filename), phaseVals)
}

func Part2() int {
	phaseVals := []int{5, 6, 7, 8, 9}
	return puzzle(input(filename), phaseVals)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), ",")
	return output
}

func puzzle(input []string, phaseVals []int) int {
	var program []int
	for i := 0; i < len(input); i++ {
		tmp, _ := strconv.Atoi(input[i])
		program = append(program, tmp)
	}
	// generate list of phase sequence
	var phasesList [][]int
	maxOutput := 0
	for _, perm := range permutations(phaseVals) {
		phasesList = append(phasesList, perm)
	}
	// Simulate the output of each phase sequence
	for _, phases := range phasesList {
		amplifier := []intcode.Computer{}
		for i, phase := range phases {
			amplifier = append(amplifier, intcode.NewComputer(&program))
			go amplifier[i].Run()
			amplifier[i].Input <- phase
			if i == 0 {
				amplifier[i].Input <- 0
			}
		}
	feedback:
		for {
			for i := range phases {
				if output, openChannel := <-amplifier[i].Output; openChannel {
					if output > maxOutput {
						maxOutput = output
					}
					if i != 4 {
						amplifier[i+1].Input <- output
					} else {
						amplifier[0].Input <- output
					}
				} else {
					break feedback
				}
			}
		}
	}
	return maxOutput
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
