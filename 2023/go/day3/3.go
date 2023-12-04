package day3

import (
	"image"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	log.SetLevel(log.DebugLevel)
	inputSlice := strings.Split(strings.TrimSpace(string(*input)), "\n")
	left, right := image.Point{X: 0, Y: -1}, image.Point{X: 0, Y: 1}
	deltas := [8]image.Point{
		right,
		left,
		{X: 1, Y: -1},  // down & left
		{X: 1, Y: 0},   // down
		{X: 1, Y: 1},   // down & right
		{X: -1, Y: -1}, // up & left
		{X: -1, Y: 0},  // up
		{X: -1, Y: 1},  // up & right
	}
	schematic := map[image.Point]interface{}{}
	X, Y := 0, len(inputSlice)
	for x, line := range inputSlice {
		X = len(line)
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(line, -1)
		indices := re.FindAllStringIndex(line, -1)
		for i, numberStr := range numbers {
			index := indices[i]
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				log.Error(err)
			}
			for i := index[0]; i < index[1]; i++ {
				schematic[image.Point{x, i}] = number
			}
		}
		re = regexp.MustCompile("[^0-9.]")
		symbols := re.FindAllString(line, -1)
		indices = re.FindAllStringIndex(line, -1)
		for i, symbol := range symbols {
			index := indices[i]
			for i := index[0]; i < index[1]; i++ {
				schematic[image.Point{x, i}] = symbol
			}
		}
	}
	sum := 0
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			k := image.Point{x, y}
			if v, ok := schematic[k]; ok && reflect.TypeOf(v).String() == "string" {
				if part2 && v.(string) != "*" {
					continue
				}
				adjacent := map[int]interface{}{}
				for _, delta := range deltas {
					d := k.Add(delta)
					if pn, ok := schematic[d]; ok {
						adjacent[pn.(int)] = nil
					}
				}
				if part2 {
					if len(adjacent) == 2 {
						adj := []int{}
						for k, _ := range adjacent {
							adj = append(adj, k)
						}
						sum += adj[0] * adj[1]
					}
				} else {
					for k, _ := range adjacent {
						sum += k
					}
				}
			}
		}
	}
	return sum
}
