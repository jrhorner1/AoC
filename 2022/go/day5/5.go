package day5

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type runeSlice []rune

func (c *runeSlice) push(r rune) { *c = append(*c, r) }
func (c *runeSlice) pop() rune   { r := (*c)[len(*c)-1]; *c = (*c)[:len(*c)-1]; return r }
func (c *runeSlice) last() rune  { return (*c)[len(*c)-1] }

func Puzzle(input *[]byte, part2 bool) string {
	//logrus.SetLevel(logrus.DebugLevel)

	stacks, instructions := parseInput(input)
	for _, instruction := range *instructions {
		if instruction == "" {
			continue
		}
		var qty, src, dest int
		_, err := fmt.Sscanf(instruction, "move %d from %d to %d", &qty, &src, &dest)
		if err != nil {
			logrus.Error(err)
		}
		if part2 {
			craneMover9001(stacks, &qty, &src, &dest)
		} else {
			moveCrates(stacks, &qty, &src, &dest)
		}
	}
	logrus.Debug(*stacks)
	printStacks(stacks)
	topCrates := []rune{}
	for _, stack := range *stacks {
		if len(stack) == 0 {
			continue
		}
		topCrates = append(topCrates, stack.last())
	}
	logrus.Debug(string(topCrates))
	return strings.ReplaceAll(string(topCrates), " ", "")
}

func parseInput(input *[]byte) (*[9]runeSlice, *[]string) {
	lines := strings.Split(string(*input), "\n")
	crates, instructions := []string{}, []string{}
	for i, line := range lines {
		if line == "" {
			crates = lines[:i]
			instructions = lines[i+1:]
			break
		}
	}
	stacks := [9]runeSlice{}
	for i := len(crates) - 2; i >= 0; i-- {
		trim := ""
		for j := 1; j < len(crates[i]); j += 4 {
			trim = trim + string(crates[i][j])
		}
		crates[i] = trim
		for j, char := range crates[i] {
			if char == ' ' {
				continue
			}
			stacks[j].push(char)
		}
	}
	logrus.Debug(stacks)
	printStacks(&stacks)
	return &stacks, &instructions
}

func moveCrates(stacks *[9]runeSlice, qty, src, dest *int) {
	for i := 0; i < *qty; i++ {
		logrus.Debug("Moving ", string((*stacks)[*src-1].last()), " from ", *src, " to ", *dest)
		(*stacks)[*dest-1].push((*stacks)[*src-1].pop())
		logrus.Debug("Moved ", string((*stacks)[*dest-1].last()), " from ", *src, " to ", *dest)
	}
}

func craneMover9001(stacks *[9]runeSlice, qty, src, dest *int) {
	temp := runeSlice{}
	for i := 0; i < *qty; i++ {
		temp.push((*stacks)[*src-1].pop())
	}
	for i := 0; i < *qty; i++ {
		(*stacks)[*dest-1].push(temp.pop())
	}
}

func printStacks(stacks *[9]runeSlice) {
	for i := 8; i >= 0; i-- {
		row := ""
		for _, stack := range *stacks {
			if len(stack)-1 < i {
				row += " "
				continue
			}
			row += string(stack[i])
		}
		logrus.Debug(row)
	}
}
