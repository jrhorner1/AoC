package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2020/input/9")
	tmp := strings.Split(strings.TrimSpace(string(input)), "\n")
	xmas := []int{}
	for _, s := range tmp {
		v, _ := strconv.Atoi(s)
		xmas = append(xmas, v)
	}
	valid := map[int]bool{}
	var invalid int
	for i, data := range xmas {
		if i < 25 {
			continue
		}
		preamble := xmas[i-25 : i]
		for j, x := range preamble {
			for _, y := range preamble[j:] {
				if x+y == data {
					valid[data] = true
				}
			}
		}
		if !valid[data] {
			invalid = data
			fmt.Println("Part 1:", data)
			break
		}
	}
	for low, x := range xmas {
		total := x
		for high, y := range xmas {
			if high <= low {
				continue
			}
			total += y
			if total > invalid {
				break
			} else if total == invalid {
				sort.Ints(xmas[low : high+1])
				fmt.Println("Part 2:", xmas[low]+xmas[high])
				return
			}
		}
	}
}
