package day3

import (
	"math"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n")
	var bitSums []int
	for i := 0; i < len(in[0]); i++ { // loop through each column of bits
		bitSums = append(bitSums, 0)
		for _, row := range in {
			// convert the bit to its ascii value and subtract the ascii value of 0 to get a integer 1 or 0
			bitSums[i] += int(row[i]) - int('0')
		}
	}
	var gammaRate, epsilonRate int
	for i, sum := range bitSums {
		if sum > len(in)/2 { // if the bit sum is greater than half of all rows, '1' is the most common bit
			// use base 2 positional notation to convert from binary to decimal
			// https://en.wikipedia.org/wiki/Binary_number#Decimal
			gammaRate += int(math.Pow(2, float64(len(bitSums)-1-i)))
		} else { // if the bit sum is less than or equal to half of all rows, '0' is the most common bit
			epsilonRate += int(math.Pow(2, float64(len(bitSums)-1-i)))
		}
	}
	if part2 {
		oxyGenRating := getRating(in, true, 0)
		co2ScrubRating := getRating(in, false, 0)
		return oxyGenRating * co2ScrubRating
	}
	return gammaRate * epsilonRate
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
