package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2021/input/1"

func Part1() int {
	return DepthScan(Input(filename), false)
}

func Part2() int {
	return DepthScan(Input(filename), true)
}

func Input(file string) []int {
	input, _ := ioutil.ReadFile(file)
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	// convert strings to ints
	var depths []int
	for i := 0; i < len(in); i++ {
		tmp, _ := strconv.Atoi(in[i])
		depths = append(depths, tmp)
	}
	return depths
}

func DepthScan(depths []int, sliding bool) int {
	increase := 0
	if sliding {
		// add 3 depths (sliding window) and compare to previous sum, counting increases
		prevSum := 0
		for i := 0; i < len(depths); i++ {
			if i < 3 {
				prevSum += depths[i]
				continue
			}
			sum := depths[i-2] + depths[i-1] + depths[i]
			if sum > prevSum {
				increase++
			}
			prevSum = sum
		}
	} else {
		// compare each depth to the previous and count the increases
		for i := 1; i < len(depths); i++ {
			if depths[i] > depths[i-1] {
				increase++
			}
		}
	}
	return increase
}
