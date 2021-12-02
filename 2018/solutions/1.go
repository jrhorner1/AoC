package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2018/input/1")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	frequency := 0
	changeList := []int{}
	for _, change := range in {
		direction := change[0]
		amount, _ := strconv.Atoi(change[1:])
		switch direction {
		case '+':
			frequency += amount
			changeList = append(changeList, amount)
		case '-':
			frequency -= amount
			changeList = append(changeList, -amount)
		}
	}
	fmt.Println("Part 1:", frequency)
	frequencyLog := []int{0}
	frequency = 0
loop:
	for {
		for _, change := range changeList {
			frequency += change
			for _, log := range frequencyLog {
				if frequency == log {
					break loop
				}
			}
			frequencyLog = append(frequencyLog, frequency)
		}
	}
	fmt.Println("Part 2:", frequency)
}
