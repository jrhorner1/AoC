package day7

import (
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

func Puzzle(input *[]byte, phaseVals []int) int {
	var program []int
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), ",") {
		tmp, _ := strconv.Atoi(i)
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
