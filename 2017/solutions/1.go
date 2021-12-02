package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2017/input/1")
	in := strings.TrimSpace(string(input))
	digits := []int{}
	// convert all of the digits into integers to make it easier to work with
	for _, d := range in {
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
	fmt.Println("Part 1:", adjacentSum)
	fmt.Println("Part 2:", halfwaySum)
}
