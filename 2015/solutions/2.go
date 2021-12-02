package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2015/input/2")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	wrappingPaper, ribbon := 0, 0
	for _, d := range in {
		dimensions := strings.Split(d, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		a := 2*l*w + 2*w*h + 2*h*l
		s := []int{l * w, w * h, h * l}
		sort.Ints(s)
		p := []int{2*l + 2*w, 2*w + 2*h, 2*h + 2*l}
		sort.Ints(p)
		v := l * w * h

		wrappingPaper += a + s[0]
		ribbon += p[0] + v
	}
	fmt.Println("Part 1:", wrappingPaper)
	fmt.Println("Part 2:", ribbon)
}
