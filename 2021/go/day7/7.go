package day7

import (
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

var filename = "2021/input/7"

// var filename = "2021/examples/7"

func Part1() float64 {
	return puzzle(input(filename), false)
}

func Part2() float64 {
	return puzzle(input(filename), true)
}

func input(file string) []float64 {
	input, _ := ioutil.ReadFile(file)
	in := strings.Split(strings.TrimSpace(string(input)), ",")
	output := []float64{}
	for _, i := range in {
		out, _ := strconv.Atoi(i)
		output = append(output, float64(out))
	}
	return output
}

func puzzle(input []float64, part2 bool) float64 {
	if part2 {
		meanFloor, meanCeil := math.Floor(Mean(&input)), math.Ceil(Mean(&input))
		fuelFloor, fuelCeil := 0.0, 0.0
		for _, crab := range input {
			fuelFloor += getSteps(crab, meanFloor)
			fuelCeil += getSteps(crab, meanCeil)
		}
		return math.Min(fuelFloor, fuelCeil)
	}
	sort.Float64s(input)
	var median, fuel float64
	if len(input)%2 != 0 { // odd
		median = input[len(input)/2-1]
	} else { // even
		index := (len(input)/2 - 1 + len(input)/2) / 2
		median = input[index]
	}
	for _, i := range input {
		fuel += math.Abs(i - median)
	}
	return fuel
}

func Mean(slice *[]float64) float64 {
	sum := 0.0
	for _, i := range *slice {
		sum += i
	}
	return sum / float64(len(*slice))
}

func getSteps(crab float64, mean float64) float64 {
	steps := math.Abs(float64(crab) - mean)
	// https://en.wikipedia.org/wiki/Triangular_number
	return math.Floor((math.Pow(steps, 2) + steps) / 2)
}
