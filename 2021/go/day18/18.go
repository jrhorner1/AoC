package day18

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
)

type Snailfish struct {
	value, depth int
}

type SnailfishSlice []Snailfish

func Puzzle(input *[]byte, part2 bool) int {
	numbers := parseInput(input)
	if part2 {
		mags := []int{}
		for _, a := range numbers {
			for _, b := range numbers {
				if reflect.DeepEqual(a, b) {
					continue
				}
				a.Add(&b)
				mags = append(mags, a.Magnitude())
				fmt.Println(len(mags), mags)
			}
		}
		sort.Ints(mags)
		return mags[len(mags)-1]
	}
	sum := SnailfishSlice{}
	for i, number := range numbers {
		if i == 0 {
			sum = number
			continue
		}
		sum.Add(&number)
	}
	return sum.Magnitude()
}

func parseInput(input *[]byte) []SnailfishSlice {
	var numbers []SnailfishSlice
	for i, line := range bytes.Split(bytes.TrimSpace(*input), []byte("\n")) {
		if i == len(*input)-1 {
			break
		}
		depth := 0
		var number SnailfishSlice
		for j := 0; j < len(line); j++ {
			switch line[j] {
			case byte('['):
				depth += 1
			case byte(','):
				continue
			case byte(']'):
				depth -= 1
			default:
				digit := map[byte]interface{}{48: nil, 49: nil, 50: nil, 51: nil, 52: nil, 53: nil, 54: nil, 55: nil, 56: nil, 57: nil}
				b := string(line[j])
				if _, ok := digit[line[j+1]]; ok {
					j += 1
					b += string(line[j])
				}
				value, _ := strconv.Atoi(b)
				number = append(number, Snailfish{value: value, depth: depth})
			}
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func (a *SnailfishSlice) Add(b *SnailfishSlice) {
	*a = append(*a, *b...)
	for i := range *a {
		(*a)[i].depth++
	}
	a.Reduce()
}

func (s *SnailfishSlice) Reduce() {
	// fmt.Println("Reducing", *s)
	for s.Explode() || s.Split() {
		// fmt.Println(*s)
	}
}

/*
*  If any pair is nested inside four pairs, the leftmost such pair explodes.
*
*  To explode a pair, the pair's left value is added to the first regular number to the left of the exploding pair (if any), and
*  the pair's right value is added to the first regular number to the right of the exploding pair (if any). Exploding pairs will
*  always consist of two regular numbers. Then, the entire exploding pair is replaced with the regular number 0.
 */
func (s *SnailfishSlice) Explode() bool {
	for i := range *s {
		if (*s)[i].depth > 4 {
			// fmt.Println("Boom!")
			if i > 0 {
				(*s)[i-1].value += (*s)[i].value
			}
			(*s)[i] = Snailfish{0, (*s)[i].depth - 1}
			if i < len(*s)-2 {
				(*s)[i+2].value += (*s)[i+1].value
				copy((*s)[i+1:], (*s)[i+2:])
			}
			(*s)[len(*s)-1] = Snailfish{}
			(*s) = (*s)[:len(*s)-1]
			return true
		}
	}
	return false
}

/*
*  If any regular number is 10 or greater, the leftmost such regular number splits.
*
*  To split a regular number, replace it with a pair; the left element of the pair should be the regular number divided by two
*  and rounded down, while the right element of the pair should be the regular number divided by two and rounded up.
 */
func (s *SnailfishSlice) Split() bool {
	for i := range *s {
		if (*s)[i].value >= 10 {
			// fmt.Println("Crack!")
			value, depth := float64((*s)[i].value)/2, (*s)[i].depth+1
			left, right := Snailfish{int(math.Floor(value)), depth}, Snailfish{int(math.Ceil(value)), depth}
			(*s)[i] = left
			(*s) = append((*s), right)
			for j := len(*s) - 1; j-1 > i; j-- {
				(*s)[j], (*s)[j-1] = (*s)[j-1], (*s)[j]
			}
			return true
		}
	}
	return false
}

/*
*  The magnitude of a pair is 3 times the magnitude of its left element plus 2 times the magnitude of its right element. The
*  magnitude of a regular number is just that number.
 */
func (s *SnailfishSlice) Magnitude() int {
	fmt.Println("Calculating magnitude ", *s)
	f := make(SnailfishSlice, len(*s))
	copy(f, *s)
loop:
	for len(f) > 1 {
		for i := 0; i < len(f)-1; i++ {
			if i == len(f)-1 {
				continue loop
			}
			if f[i+1].depth == f[i].depth {
				magnitude := (f[i].value * 3) + (f[i+1].value * 2)
				f[i] = Snailfish{value: magnitude, depth: f[i].depth - 1}
				if i < len(f)-2 {
					copy(f[i+1:], f[i+2:])
				}
				f[len(f)-1] = Snailfish{}
				f = f[:len(f)-1]
			}
		}
	}
	return f[0].value
}
