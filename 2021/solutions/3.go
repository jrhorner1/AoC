package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	diag := parse("2021/input/3")
	bits := len(diag[0])

	gamma, epsilon := "", ""
	for i := 0; i < bits; i++ {
		bit0 := 0
		for _, row := range diag {
			if row[i] == 0 {
				bit0++
			}
		}
		if bit0 >= len(diag)/2 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammaRate, _ := strconv.ParseUint(gamma, 2, 32)
	epsilonRate, _ := strconv.ParseUint(epsilon, 2, 32)
	powerConsumption := gammaRate * epsilonRate
	fmt.Println("Part 1:", powerConsumption)

	oxyGenRating := getRating(diag, true, 0)
	co2ScrubRating := getRating(diag, false, 0)
	lifeSupportRating := oxyGenRating * co2ScrubRating
	fmt.Println("Part 2:", lifeSupportRating)
	fmt.Println("Happy Holidays 2021!")
}

func parse(file string) [][]int {
	input, _ := ioutil.ReadFile(file)
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	diag := [][]int{}
	for _, s := range in {
		row := []int{}
		for _, r := range s {
			switch r {
			case '0':
				row = append(row, 0)
			case '1':
				row = append(row, 1)
			}
		}
		diag = append(diag, row)
	}
	return diag
}

func getRating(diagData [][]int, majority bool, index int) int {
	// base case
	if len(diagData) == 1 {
		binaryString := parseBin(diagData[0])
		rating, _ := strconv.ParseInt(binaryString, 2, 32)
		return int(rating)
	}
	// divide the data into 2 slices based on bit value at index
	var zerosData, onesData [][]int
	for _, binarySlice := range diagData {
		if binarySlice[index] == 0 {
			zerosData = append(zerosData, binarySlice)
		} else {
			onesData = append(onesData, binarySlice)
		}
	}
	// recurse with the data set that matches the bit criteria
	if len(zerosData) > len(onesData) == majority {
		return getRating(zerosData, majority, index+1)
	}
	return getRating(onesData, majority, index+1)
}

func parseBin(binarySlice []int) string {
	binaryString := ""
	for _, bit := range binarySlice {
		binaryString += strconv.Itoa(bit)
	}
	return binaryString
}
