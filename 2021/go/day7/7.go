package day7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) float64 {
	in := strings.Split(strings.TrimSpace(string(*input)), ",")
	crabs := []float64{}
	for _, i := range in {
		out, _ := strconv.Atoi(i)
		crabs = append(crabs, float64(out))
	}
	if part2 {
		meanFloor, meanCeil := math.Floor(Mean(&crabs)), math.Ceil(Mean(&crabs))
		fuelFloor, fuelCeil := 0.0, 0.0
		for _, crab := range crabs {
			fuelFloor += getSteps(crab, meanFloor)
			fuelCeil += getSteps(crab, meanCeil)
		}
		return math.Min(fuelFloor, fuelCeil)
	}
	sort.Float64s(crabs)
	var median, fuel float64
	if len(crabs)%2 != 0 { // odd
		median = crabs[len(crabs)/2-1]
	} else { // even
		index := (len(crabs)/2 - 1 + len(crabs)/2) / 2
		median = crabs[index]
	}
	for _, i := range crabs {
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
