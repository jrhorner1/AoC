package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

/*
    0  1  2  3  4
    5  6  7  8  9
   10 11 12 13 14
   15 16 17 18 19
   20 21 22 23 24
*/

type bingoSquare struct {
	number int
	marked bool
}

func main() {
	input, _ := ioutil.ReadFile("2021/input/4")
	in := strings.Split(strings.TrimSpace(string(input)), "\n\n") // split on double newline to get each board into a single string
	numbers := parseNumbers(in[0])
	boards := parseBoards(in[1:])
	scores := []int{}
	won := make(map[int]int)
	for c, called := range numbers {
		for i := range boards {
			for j := range boards[i] {
				if boards[i][j].number == called {
					boards[i][j].marked = true
				}
			}
			if c > int(math.Sqrt(float64(len(boards[i])))) { // check if at least 5 numbers have been called already
				if _, yes := won[i]; !yes { // check if the board in question has't won already and that
					if checkBingo(&boards[i]) {
						score := getScore(&boards[i], called)
						scores = append(scores, score)
						won[i] = score
					}
				}
			}
		}
	}
	fmt.Println("Part 1:", scores[0])
	fmt.Println("Part 2:", scores[len(scores)-1])
	fmt.Println("Happy Holidays 2021!")
}

func parseNumbers(input string) []int {
	splitStrings := strings.Split(input, ",")
	numbers := []int{}
	for _, s := range splitStrings {
		number, _ := strconv.Atoi(s)
		numbers = append(numbers, number)
	}
	return numbers
}

func parseBoards(input []string) [][]bingoSquare {
	boards := [][]bingoSquare{}
	for _, boardString := range input {
		board := []bingoSquare{}
		for _, rowString := range strings.Split(strings.TrimSpace(string(boardString)), "\n") {
			numberStrings := strings.Fields(rowString)
			for _, numberString := range numberStrings {
				number, _ := strconv.Atoi(numberString)
				space := bingoSquare{number, false}
				board = append(board, space)
			}
		}
		boards = append(boards, board)
	}
	return boards
}

func checkRowBingo(board *[]bingoSquare) bool {
	sqTotal := len(*board)
	sqPerSide := int(math.Sqrt(float64(sqTotal)))
	tally := 0
	for i := 0; i < sqTotal; i += sqPerSide {
		for j := i; j <= i+(sqPerSide-1); j++ {
			if (*board)[j].marked {
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

func checkColumnBingo(board *[]bingoSquare) bool {
	sqTotal := len(*board)
	sqPerSide := int(math.Sqrt(float64(sqTotal)))
	tally := 0
	for i := 0; i < sqPerSide; i++ {
		for j := i; j <= i+(sqPerSide*(sqPerSide-1)); j += sqPerSide {
			if (*board)[j].marked {
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

func checkBingo(board *[]bingoSquare) bool {
	if checkRowBingo(board) {
		return true
	} else if checkColumnBingo(board) {
		return true
	} else {
		return false
	}
}

func getScore(board *[]bingoSquare, lastCalled int) int {
	unmarkedSum := 0
	for _, space := range *board {
		if !space.marked {
			unmarkedSum += space.number
		}
	}
	return unmarkedSum * lastCalled
}
