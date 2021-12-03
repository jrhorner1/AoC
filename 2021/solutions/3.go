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

	oxyGenRating := getRating(diag, 0, 0)
	co2ScrubRating := getRating(diag, 1, 0)
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

func getRating(diagData [][]int, bitCriteria int, index int) int {
	// count bit frequency
	zeros, ones := 0, 0
	for _, binarySlice := range diagData {
		switch binarySlice[index] {
		case 0:
			zeros++
		case 1:
			ones++
		}
	}
	//determine majority bit
	var majority int
	var minority int
	if zeros > ones {
		majority, minority = 0, 1
	} else {
		majority, minority = 1, 0
	}
	// create a new slice based on bit criteria
	var newData [][]int
	switch bitCriteria {
	case 0:
		for _, binarySlice := range diagData {
			if binarySlice[index] == majority {
				newData = append(newData, binarySlice)
			}
		}
	case 1:
		for _, binarySlice := range diagData {
			if binarySlice[index] == minority {
				newData = append(newData, binarySlice)
			}
		}
	}
	// recurse
	if len(newData) != 1 {
		return getRating(newData, bitCriteria, index+1)
	}
	// convert the integer slice to a binary string then to a decimal number
	binaryString := parseBin(newData[0])
	rating, _ := strconv.ParseUint(binaryString, 2, 32)
	return int(rating)
}

func parseBin(binarySlice []int) string {
	binaryString := ""
	for _, bit := range binarySlice {
		binaryString += strconv.Itoa(bit)
	}
	return binaryString
}
