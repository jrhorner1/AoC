package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "2017/input/1"

func Part1() int {
	return Puzzle(Input(filename), false)
}

func Part2() int {
	return Puzzle(Input(filename), true)
}

func Input(file string) string {
	input, _ := ioutil.ReadFile(file)
	output := strings.TrimSpace(string(input))
	return output
}

func Puzzle(input string, part2 bool) int {
	digits := []int{}
	// convert all of the digits into integers to make it easier to work with
	for _, d := range input {
		digit, _ := strconv.Atoi(string(d))
		digits = append(digits, digit)
	}
	adjacentSum, halfwaySum := 0, 0
	for i, digit := range digits {
		nextDigit, halfDigit := 0, 0
		if i == len(digits)-1 {
			nextDigit = digits[0]
		} else {
			nextDigit = digits[i+1]
		}
		if digit == nextDigit {
			adjacentSum += digit
		}
		if i <= (len(digits)/2)-1 { // only need to test half of the digits
			halfDigit = digits[i+(len(digits)/2)]
			if digit == halfDigit {
				halfwaySum += digit * 2 // multiply for the duplicate in the other half
			}
		}
	}
	if part2 {
		return halfwaySum
	}
	return adjacentSum
}
