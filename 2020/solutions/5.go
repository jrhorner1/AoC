package main

import (
	"fmt"
	"sort"
	"strconv"

	"../utils"
)

type BoardingPass struct {
	Row int64
	Col int64
	Id  int64
}

func parsebin(i string) string {
	var bin string
	for _, j := range i {
		switch string(j) {
		case "F", "L":
			bin += "0"
		case "B", "R":
			bin += "1"
		}
	}
	return bin
}

func findseat(bin string) BoardingPass {
	var bp BoardingPass
	bp.Row, _ = strconv.ParseInt(bin[:7], 2, 64)
	bp.Col, _ = strconv.ParseInt(bin[7:], 2, 64)
	bp.Id = (bp.Row * 8) + bp.Col
	return bp
}

type ByRowCol []BoardingPass

func (a ByRowCol) Len() int      { return len(a) }
func (a ByRowCol) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRowCol) Less(i, j int) bool {
	if a[i].Row < a[j].Row {
		return true
	}
	if a[i].Row > a[j].Row {
		return false
	}
	return a[i].Col < a[j].Col
}

func main() {
	input := utils.OpenFile()
	var bps []BoardingPass
	for _, i := range input {
		bin := parsebin(i)
		pass := findseat(bin)
		bps = append(bps, pass)
	}
	sort.Sort(ByRowCol(bps))
	fmt.Println("Part 1:", bps[len(bps)-1].Id)
	var my_bp BoardingPass
	for i, bp := range bps {
		if bps[i+1].Id != bp.Id+1 && bps[i+1].Id == bp.Id+2 {
			my_bp.Id = bp.Id + 1
			my_bp.Row = bp.Row
			my_bp.Col = bp.Col + 1
			break
		}
	}
	fmt.Println("Part 2:", my_bp.Id)
}
