package day1

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n")
	fuel := 0
	for i := 0; i < len(in); i++ {
		mass, _ := strconv.Atoi(in[i])
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
