package day10

import (
	"sort"
	"strings"
)

type runeSlice []rune

func (c *runeSlice) push(r rune) { *c = append(*c, r) }
func (c *runeSlice) pop() rune   { r := (*c)[len(*c)-1]; *c = (*c)[:len(*c)-1]; return r }
func (c *runeSlice) last() rune  { return (*c)[len(*c)-1] }

func Puzzle(input *[]byte, part2 bool) int {
	completions := []runeSlice{}
	illegal := runeSlice{}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		stack := runeSlice{}
		discard := runeSlice{}
		completion := runeSlice{}
		for _, r := range line {
			switch r {
			case ')', ']', '}', '>':
				chunks := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
				if stack.last() == chunks[r] {
					stack.pop()
				} else {
					discard.push(r)
				}
			default:
				stack.push(r)
			}
		}
		if len(discard) > 0 {
			illegal.push(discard[0])
		} else if len(stack) > 0 {
			chunks := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
			for len(stack) > 0 {
				completion.push(chunks[stack.pop()])
			}
			completions = append(completions, completion)
		}
	}
	if part2 {
		scores := []int{}
		points := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
		for _, completion := range completions {
			score := 0
			for _, r := range completion {
				score *= 5
				score += points[r]
			}
			scores = append(scores, score)
		}
		sort.Ints(scores)
		return scores[(len(scores) / 2)]
	}
	score := 0
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	for _, r := range illegal {
		score += points[r]
	}
	return score
}
