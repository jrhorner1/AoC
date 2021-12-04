package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type bingoSquare struct {
	number int
	marked bool
}

type bingoBoard struct {
	squares []bingoSquare
	id      int
	rank    int
	score   int
}

func main() {
	input, _ := ioutil.ReadFile("2021/input/4")
	// input, _ := ioutil.ReadFile("2021/examples/4")
	in := strings.Split(strings.TrimSpace(string(input)), "\n\n") // split on double newline to get each board into a single string
	splitStrings := strings.Split(in[0], ",")
	numbers := []int{}
	for _, s := range splitStrings {
		number, _ := strconv.Atoi(s)
		numbers = append(numbers, number)
	}
	boards := []bingoBoard{}
	for id, boardString := range in[1:] {
		board := bingoBoard{[]bingoSquare{}, id, 0, 0}
		for _, rowString := range strings.Split(strings.TrimSpace(string(boardString)), "\n") {
			numberStrings := strings.Fields(rowString)
			for _, numberString := range numberStrings {
				number, _ := strconv.Atoi(numberString)
				space := bingoSquare{number, false}
				board.squares = append(board.squares, space)
			}
		}
		boards = append(boards, board)
	}

	ranking := 1
	for _, number := range numbers {
		for i := range boards {
			for j := range boards[i].squares {
				if boards[i].squares[j].number == number {
					boards[i].squares[j].marked = true
				}
			}
			if boards[i].rank == 0 { // check if the board in question has't won already
				if checkRowBingo(&boards[i]) || checkColumnBingo(&boards[i]) {
					unmarkedSum := 0
					for j := range boards[i].squares {
						if !boards[i].squares[j].marked {
							unmarkedSum += boards[i].squares[j].number
						}
					}
					boards[i].score = unmarkedSum * number
					boards[i].rank = ranking
					ranking++
				}
			}
		}
	}
	var first, last bingoBoard
	for i := range boards {
		if boards[i].rank == 1 {
			first = boards[i]
		}
		if boards[i].rank == len(boards) {
			last = boards[i]
		}
	}
	fmt.Println("Part 1:", first.score)
	fmt.Println("Part 2:", last.score)
	fmt.Println("Happy Holidays 2021!")
}

func checkRowBingo(board *bingoBoard) bool {
	sqTotal := len((*board).squares)
	sqPerSide := int(math.Sqrt(float64(sqTotal)))
	tally := 0
	for i := 0; i < sqTotal; i += sqPerSide {
		for j := i; j <= i+(sqPerSide-1); j++ {
			if (*board).squares[j].marked {
				tally += 1
			} else {
				break
			}
		}
		if tally == sqPerSide {
			return true
		}
		tally = 0
	}
	return false
}

func checkColumnBingo(board *bingoBoard) bool {
	sqTotal := len((*board).squares)
	sqPerSide := int(math.Sqrt(float64(sqTotal)))
	tally := 0
	for i := 0; i < sqPerSide; i++ {
		for j := i; j <= i+(sqPerSide*(sqPerSide-1)); j += sqPerSide {
			if (*board).squares[j].marked {
				tally += 1
			} else {
				break
			}
		}
		if tally == sqPerSide {
			return true
		}
		tally = 0
	}
	return false
}
