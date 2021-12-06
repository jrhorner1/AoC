package day6

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2021/input/6"

// var filename = "2021/examples/6"

func Part1() uint {
	return Puzzle(Input(filename), 80)
}

func Part2() uint {
	return Puzzle(Input(filename), 256)
}

func Input(file string) map[int]uint {
	input, _ := ioutil.ReadFile(file)
	in := strings.Split(strings.TrimSpace(string(input)), ",")
	output := map[int]uint{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, i := range in {
		out, _ := strconv.Atoi(i)
		output[out] += 1
	}
	return output
}

func Puzzle(input map[int]uint, days int) uint {
	total := uint(0)
	for day := 0; day < days; day++ {
		tmp := input[0]
		for i := 0; i < 8; i++ {
			input[i] = input[i+1]
		}
		input[6] += tmp
		input[8] = tmp
	}
	for _, v := range input {
		total += v
	}
	return total
}
