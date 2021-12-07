package day5

import (
	"sort"
	"strconv"
	"strings"
)

type BoardingPass struct {
	Row int
	Col int
	Id  int
}

func Puzzle(input *[]byte, part2 bool) int {
	var bps []BoardingPass
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		bin := parsebin(i)
		pass := findseat(bin)
		bps = append(bps, pass)
	}
	sort.Sort(ByRowCol(bps))
	if part2 {
		var my_bp BoardingPass
		for i, bp := range bps {
			if bps[i+1].Id != bp.Id+1 && bps[i+1].Id == bp.Id+2 {
				my_bp.Id = bp.Id + 1
				my_bp.Row = bp.Row
				my_bp.Col = bp.Col + 1
				break
			}
		}
		return my_bp.Id
	}
	return bps[len(bps)-1].Id
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

func binaryStringtoInt(s string) int {
	out, _ := strconv.ParseInt(s, 2, 0)
	return int(out)
}

func findseat(bin string) BoardingPass {
	var bp BoardingPass
	bp.Row = binaryStringtoInt(bin[:7])
	bp.Col = binaryStringtoInt(bin[7:])
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
