package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fuelRequired(mass int) int {
	fuel_req := (mass / 3) - 2
	if fuel_req < 0 {
		return 0
	}
	return fuel_req
}

func main() {
	// change to request the input file from user
	input, _ := ioutil.ReadFile("2019/input/1")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")

	fuel1, fuel2 := 0, 0
	for i := 0; i < len(in); i++ {
		mass, _ := strconv.Atoi(in[i])
		fuel1 += (mass / 3) - 2
		fuelReq := fuelRequired(mass)
		for fuelReq > 0 {
			fuel2 += fuelReq
			fuelReq = fuelRequired(fuelReq)
		}
	}
	fmt.Println("Part 1:", fuel1)
	fmt.Println("Part 2:", fuel2)
}
