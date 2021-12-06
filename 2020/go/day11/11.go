package day11

import (
	"image"
	"io/ioutil"
	"strings"
)

var filename = "2020/input/11"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

func puzzle(input []string, part2 bool) int {
	seats := map[image.Point]rune{}
	for y, row := range input {
		for x, rune := range row {
			seats[image.Point{x, y}] = rune
		}
	}
	if part2 {
		return sim(seats, 5, func(a, b image.Point) image.Point {
			for seats[a.Add(b)] == '.' {
				a = a.Add(b)
			}
			return a.Add(b)
		})
	}
	return sim(seats, 4, func(a, b image.Point) image.Point { return a.Add(b) })
}

func sim(seats map[image.Point]rune, spookFactor int, adjacent func(a, b image.Point) image.Point) int {
	adjacentSeats := []image.Point{
		{-1, -1}, // top left
		{-1, 0},  // top
		{-1, 1},  // top right
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // bottom left
		{1, 0},   // bottom
		{1, 1},   // bottom right
	}
	occupiedSeats, iter := 0, 0
	for loop := true; loop; iter++ {
		occupiedSeats, loop = 0, false // reset occupiedSeats for each iteration
		occupied, vacant := '#', 'L'
		new_seats := map[image.Point]rune{}
		for seat, status := range seats {
			adjacentOccupied := 0
			for _, adjacentSeat := range adjacentSeats {
				if seats[adjacent(seat, adjacentSeat)] == occupied {
					adjacentOccupied++
				}
			}
			if status == occupied && adjacentOccupied >= spookFactor {
				status = vacant
			} else if status == vacant && adjacentOccupied == 0 || status == occupied {
				status = occupied
				occupiedSeats++
			}
			new_seats[seat] = status
			loop = loop || new_seats[seat] != seats[seat]
		}
		seats = new_seats
	}
	return occupiedSeats
}
