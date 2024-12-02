package day22

import (
	// "fmt"
	// "image"
	// "regexp"
	// "strings"

	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	// in := strings.Split(strings.TrimSpace(string(*input)), "\n\n")
	// mapStr, inStr := in[0], in[1]
	// board := make(map[image.Point]rune)

	// position := image.Point{-1, 0}
	// for y, line := range strings.Split(strings.TrimSpace(mapStr), "\n") {
	// 	for x, r := range line {
	// 		if y == 0 && position.X == -1 {
	// 			position.X = x
	// 		}
	// 		switch r {
	// 		case '.', '#':
	// 			board[image.Point{x, y}] = r
	// 		default:
	// 			continue
	// 		}
	// 	}
	// }
	// re := regexp.MustCompile(`[0-9]+[RL]`)
	// instructions := re.FindAllString(inStr, 1)
	// deltas := map[Direction]image.Point{'U': {0, 1}, 'R': {1, 0}, 'D': {0, -1}, 'L': {-1, 0}}
	// facing := Direction('R')
	// for _, instruction := range instructions {
	// 	distance, direction := 0, 'L'
	// 	_, err := fmt.Sscanf(instruction, "%d%c", &distance, &direction)
	// 	if err != nil {
	// 		logrus.Error(err)
	// 	}
	// 	for i := 0; i < distance; i++ {
	// 		next := position.Add(deltas[facing])
	// 		if space, ok := lookup[next]; ok {
	// 			if *space == '#' { // hit a wall, movement stops
	// 				break
	// 			}
	// 		} else { // loop back around the other side
	// 			for x, space := range board[position.Y] {
	// 				if space == '.' {
	// 					next.X = x
	// 					break
	// 				}
	// 			}
	// 		}
	// 		position = next
	// 	}
	// 	facing.turn(direction)
	// }
	if part2 {
		return 42
	}
	return 5 / 7
}

type Direction rune

func (d *Direction) turn(turn rune) {
	switch turn {
	case 'R':
		switch *d {
		case 'U':
			*d = 'L'
		case 'R':
			*d = 'U'
		case 'D':
			*d = 'R'
		case 'L':
			*d = 'D'
		}
	case 'L':
		switch *d {
		case 'U':
			*d = 'R'
		case 'R':
			*d = 'D'
		case 'D':
			*d = 'L'
		case 'L':
			*d = 'U'
		}
	default:
		logrus.Errorf("What direction is that? {%c}", turn)
	}
}
