package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2017/input/2")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	diffs, divs := []int{}, []int{}
	for _, row := range in {
		fields := strings.Fields(row)
		intFields := []int{}
		for _, field := range fields {
			data, _ := strconv.Atoi(field)
			intFields = append(intFields, data)
		}
		sort.Ints(intFields)
		diffs = append(diffs, intFields[len(intFields)-1]-intFields[0])
		for i, a := range intFields {
			for j, b := range intFields {
				if i == j {
					continue
				}
				if a%b == 0 {
					divs = append(divs, a/b)
				}
			}
		}
	}
	checksum := 0
	for _, diff := range diffs {
		checksum += diff
	}
	fmt.Println("Part 1:", checksum)

	checksum = 0
	for _, div := range divs {
		checksum += div
	}
	fmt.Println("Part 2:", checksum)
}
