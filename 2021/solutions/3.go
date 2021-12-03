package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2021/input/3")
	diag := strings.Split(strings.TrimSpace(string(input)), "\n")

	var bitSums []int
	for i := 0; i < len(diag[0]); i++ {
		bitSums = append(bitSums, 0)
		for _, bit := range diag {
			bitSums[i] += int(bit[i]) - int('0')
		}
	}
	var gammaRate, epsilonRate int
	for i, sum := range bitSums {
		if sum > len(diag)/2 {
			gammaRate += int(math.Pow(2, float64(len(bitSums)-1-i)))
		} else {
			epsilonRate += int(math.Pow(2, float64(len(bitSums)-1-i)))
		}
	}
	fmt.Println("Part 1:", gammaRate*epsilonRate)

	oxyGenRating := getRating(diag, true, 0)
	co2ScrubRating := getRating(diag, false, 0)
	fmt.Println("Part 2:", oxyGenRating*co2ScrubRating)

	fmt.Println("Happy Holidays 2021!")
}

func getRating(data []string, majority bool, index int) int {
	// base case
	if len(data) == 1 {
		rating, _ := strconv.ParseInt(data[0], 2, 32)
		return int(rating)
	}
	// divide the data into 2 slices based on bit value at index
	var zerosData, onesData []string
	for _, binaryString := range data {
		if binaryString[index] == '0' {
			zerosData = append(zerosData, binaryString)
		} else {
			onesData = append(onesData, binaryString)
		}
	}
	// recurse with the data set that matches the bit criteria
	if len(zerosData) > len(onesData) == majority {
		return getRating(zerosData, majority, index+1)
	}
	return getRating(onesData, majority, index+1)
}
