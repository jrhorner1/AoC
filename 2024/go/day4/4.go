package day4

import (
	"fmt"
	"strings"

	geo "github.com/jrhorner1/AoC/pkg/math/geometry"
	log "github.com/sirupsen/logrus"
)

var DEBUG = false

func Puzzle(input *[]byte, part2 bool) int {
	if DEBUG {
		log.SetLevel(log.DebugLevel)
	}
	ymax, xmax := 0, 0
	wordSearch := make(map[geo.Point]rune)
	for y, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		for x, r := range line {
			wordSearch[geo.Point{X: x, Y: y}] = r
			xmax = x + 1
		}
		ymax = y + 1
	}
	if log.GetLevel() == log.DebugLevel {
		printWS(&wordSearch, ymax, xmax)
	}

	found := make(map[geo.Point]rune)
	wordCount := 0

	if part2 {
		deltas := [][]geo.Point{
			// horizontal + vertical
			// {geo.Point{X: 1, Y: 0}, geo.Point{X: -1, Y: 0}, geo.Point{X: 0, Y: -1}, geo.Point{X: 0, Y: 1}}, // right left up down
			// {geo.Point{X: -1, Y: 0}, geo.Point{X: 1, Y: 0}, geo.Point{X: 0, Y: -1}, geo.Point{X: 0, Y: 1}}, // left right up down
			// {geo.Point{X: 1, Y: 0}, geo.Point{X: -1, Y: 0}, geo.Point{X: 0, Y: 1}, geo.Point{X: 0, Y: -1}}, // right left down up
			// {geo.Point{X: -1, Y: 0}, geo.Point{X: 1, Y: 0}, geo.Point{X: 0, Y: 1}, geo.Point{X: 0, Y: -1}}, // left right down up
			// diagonal
			{geo.Point{X: 1, Y: -1}, geo.Point{X: -1, Y: 1}, geo.Point{X: -1, Y: -1}, geo.Point{X: 1, Y: 1}}, // rd lu ld ru
			{geo.Point{X: -1, Y: 1}, geo.Point{X: 1, Y: -1}, geo.Point{X: -1, Y: -1}, geo.Point{X: 1, Y: 1}}, // lu rd ld ru
			{geo.Point{X: 1, Y: -1}, geo.Point{X: -1, Y: 1}, geo.Point{X: 1, Y: 1}, geo.Point{X: -1, Y: -1}}, // rd lu ru ld
			{geo.Point{X: -1, Y: 1}, geo.Point{X: 1, Y: -1}, geo.Point{X: 1, Y: 1}, geo.Point{X: -1, Y: -1}}, // lu rd ru ld
		}
		for y := 0; y < ymax; y++ {
			for x := 0; x < xmax; x++ {
				base := geo.Point{X: x, Y: y}
				if wordSearch[base] == 'A' { // if A, check deltas
					for _, delta := range deltas {
						for i, d := range delta {
							curr := base.Add(d)
							if (i == 0 || i == 2) && wordSearch[curr] == 'M' { // if M, continue
								continue
							} else if (i == 1) && wordSearch[curr] == 'S' { // if S, continue
								continue
							} else if (i == 3) && wordSearch[curr] == 'S' { // if S, continue
								wordCount++
								// record find
								found[base] = 'A'
								for i, d := range delta {
									if i == 0 || i == 2 {
										found[base.Add(d)] = 'M'
									} else if i == 1 || i == 3 {
										found[base.Add(d)] = 'S'
									}
								}
							}
							break
						}
					}
				}
			}
		}
		if log.GetLevel() == log.DebugLevel {
			printWS(&found, ymax, xmax)
		}
		return wordCount
	}

	deltas := [][]geo.Point{
		{geo.Point{X: 1, Y: 0}, geo.Point{X: 2, Y: 0}, geo.Point{X: 3, Y: 0}},       // right
		{geo.Point{X: 1, Y: -1}, geo.Point{X: 2, Y: -2}, geo.Point{X: 3, Y: -3}},    // diag (right & down)
		{geo.Point{X: 0, Y: -1}, geo.Point{X: 0, Y: -2}, geo.Point{X: 0, Y: -3}},    // down
		{geo.Point{X: -1, Y: -1}, geo.Point{X: -2, Y: -2}, geo.Point{X: -3, Y: -3}}, // diag (left & down)
		{geo.Point{X: -1, Y: 0}, geo.Point{X: -2, Y: 0}, geo.Point{X: -3, Y: 0}},    // left
		{geo.Point{X: -1, Y: 1}, geo.Point{X: -2, Y: 2}, geo.Point{X: -3, Y: 3}},    // diag (left & up)
		{geo.Point{X: 0, Y: 1}, geo.Point{X: 0, Y: 2}, geo.Point{X: 0, Y: 3}},       // up
		{geo.Point{X: 1, Y: 1}, geo.Point{X: 2, Y: 2}, geo.Point{X: 3, Y: 3}},       // diag (right & up)
	}

	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			base := geo.Point{X: x, Y: y}
			if wordSearch[base] == 'X' { // if X, check deltas
				for _, delta := range deltas {
					for i, d := range delta {
						curr := base.Add(d)
						if i == 0 && wordSearch[curr] == 'M' { // if M, continue
							continue
						} else if i == 1 && wordSearch[curr] == 'A' { // if A, continue
							continue
						} else if i == 2 && wordSearch[curr] == 'S' { // if S, increment count
							wordCount++
							// record find
							found[base] = 'X'
							for i, d := range delta {
								switch i {
								case 0:
									found[base.Add(d)] = 'M'
								case 1:
									found[base.Add(d)] = 'A'
								case 2:
									found[base.Add(d)] = 'S'
								}
							}
						}
						break
					}
				}
			}
		}
	}
	if log.GetLevel() == log.DebugLevel {
		printWS(&found, ymax, xmax)
	}
	return wordCount
}

func printWS(ws *map[geo.Point]rune, ymax, xmax int) {
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			if r, ok := (*ws)[geo.Point{X: x, Y: y}]; ok {
				fmt.Printf("%c", r)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
