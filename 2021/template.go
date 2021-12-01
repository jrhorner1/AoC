package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println("Part 1:", in)
	// fmt.Println("Part 2:", input)
	fmt.Println("Happy Holidays 2021!")
}
