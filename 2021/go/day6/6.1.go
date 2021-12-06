package day6

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Don't use this.

var example = "2021/examples/6"

func YeetRevengeOfTheYote() int {
	return yeet(input(example), 80)
}

func Yeet2ElectricBoogaloo() int {
	return yeet(input(example), 256)
}

func input(file string) []int {
	input, _ := ioutil.ReadFile(file)
	in := strings.Split(strings.TrimSpace(string(input)), ",")
	output := []int{}
	for _, i := range in {
		out, _ := strconv.Atoi(i)
		output = append(output, out)
	}
	return output
}

func yeet(input []int, days int) int {
	for i := 0; i < days; i++ {
		for j := range input {
			switch input[j] {
			case 0:
				input = append(input, 8)
				input[j] = 6
			default:
				input[j]--
			}
		}
	}
	return len(input)
}
