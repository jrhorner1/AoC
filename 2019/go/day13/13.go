package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

type Tile struct {
	x, y int
	id   int
}

func Puzzle(input *[]byte, part2 bool) int {
	var program []int
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), ",") {
		out, _ := strconv.Atoi(i)
		program = append(program, out)
	}
	// if part2 {
	// 	return arcade(&program, 2)
	// }
	computer := intcode.NewComputer(&program)
	go computer.Run()
	ok := true
	var out int
	var grid []Tile
	var tile Tile
	count, ans := 4, 0
	for ok {
		out, ok = <-computer.Output
		switch count % 3 {
		case 0:
			switch out {
			case 0:
				tile.id = out
			case 1:
				tile.id = out
			case 2:
				ans++
				tile.id = out
			case 3:
				tile.id = out
			case 4:
				tile.id = out
			}
			grid = append(grid, tile)
		case 1:
			tile.x = out
		case 2:
			tile.y = out
		}
		count++
	}
	return ans
}

func arcade(program *[]int, quarters int) int {
	computer := intcode.NewComputer(program)
	(*computer.Memory)[0] = quarters
	go computer.Run()
	ok := true
	var out int
	var grid []Tile
	var tile Tile
	var score bool = false
	count, ans := 4, 0
	input := computer.GetInput()
	for ok {
		out, ok = <-computer.GetOutput()
		switch count % 3 {
		case 0:
			if score {
				fmt.Println("Score:", out)
				score = false
				continue
			}
			switch out {
			case 0:
				tile.id = out
			case 1:
				tile.id = out
			case 2:
				ans++
				tile.id = out
			case 3:
				tile.id = out
			case 4:
				tile.id = out
			}
			grid = append(grid, tile)
		case 1:
			switch out {
			case -1:
				score = true
			default:
				tile.x = out
			}
		case 2:
			if score {
				continue
			}
			tile.y = out
		}
		count++
		input <- 0
	}
	return ans
}
