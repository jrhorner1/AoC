package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2019/input/1"

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

func puzzle(input []string, part2 bool) int {
	fuel := 0
	for i := 0; i < len(input); i++ {
		mass, _ := strconv.Atoi(input[i])
		if part2 {
			fuelReq := fuelRequired(mass)
			for fuelReq > 0 {
				fuel += fuelReq
				fuelReq = fuelRequired(fuelReq)
			}
		} else {
			fuel += (mass / 3) - 2
		}
	}
	return fuel
}

func fuelRequired(mass int) int {
	fuel_req := (mass / 3) - 2
	if fuel_req < 0 {
		return 0
	}
	return fuel_req
}
