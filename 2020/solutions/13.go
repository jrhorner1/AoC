package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("2020/input/13")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	timestamp, _ := strconv.Atoi(in[0])
	bus, wait, ts, increment := math.MaxInt64, 0, 0, 1
	for i, s := range strings.Split(in[1], ",") {
		id, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		// fmt.Println(id, "- (", timestamp, "%", id, ") =", id - timestamp%id)
		if id-timestamp%id < bus-timestamp%bus {
			bus, wait = id, id-timestamp%id
		}
		// fmt.Println("(", ts, "+", i, ") %", id, "=", (ts + i) % id)
		for (ts+i)%id != 0 { // bus id should divide cleanly into timestamp + index
			ts += increment // otherwise increment timestamp
			// fmt.Println("(", ts, "+", i, ") %", id, "=", (ts + i) % id)
		}
		increment *= id // set timestamp increment to current increment multiplied by current bus id
	}
	fmt.Println("Part 1:", bus*wait)
	fmt.Println("Part 2:", ts)
}
