package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2021/input/1")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	// convert strings to ints
	var depths []int
	for i := 0; i < len(in); i++ {
		tmp, _ := strconv.Atoi(in[i])
		depths = append(depths, tmp)
	}
	// compare each depth to the previous and count the increases
	incr1 := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			incr1++
		}
	}
	fmt.Println("Part 1:", incr1)
	// add 3 depths (sliding window) and compare to previous sum, counting increases
	incr2 := 0
	prevSum := depths[2] + depths[1] + depths[0]
	for i := 3; i < len(depths); i++ {
		sum := depths[i-2] + depths[i-1] + depths[i]
		if sum > prevSum {
			incr2++
		}
		prevSum = sum
	}
	fmt.Println("Part 2:", incr2)

	fmt.Println("Happy Holidays 2021!")
}
