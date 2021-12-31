package day24

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jrhorner1/AoC/pkg/math"
)

type State struct {
	register [4]int
	min, max uint
}

type Cache []State

func (c Cache) Len() int { return len(c) }
func (c Cache) Less(a, b int) bool {
	for i := 0; i < 4; i++ {
		if (c)[a].register[i] == (c)[b].register[i] {
			continue
		}
		return (c)[a].register[i] < (c)[b].register[i]
	}
	return false
}
func (c Cache) Swap(a, b int) { c[a], c[b] = c[b], c[a] }

func registerIndex(s string) int {
	switch s {
	case "w":
		return 0
	case "x":
		return 1
	case "y":
		return 2
	case "z":
		return 3
	}
	return -1
}

func registerValue(s string, register [4]int) int {
	switch s {
	case "w", "x", "y", "z":
		return register[registerIndex(s)]
	default:
		i, _ := strconv.Atoi(s)
		return i
	}
}

func Puzzle(input *[]byte) (max, min uint) {
	fmt.Println("Permutations:")
	cache := Cache{{[4]int{0, 0, 0, 0}, 0, 0}}
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		words := strings.Split(line, " ")
		instruction, a, b := words[0], words[1], ""
		if len(words) == 3 {
			b = words[2]
		}
		switch instruction {
		case "inp":
			for c := range cache {
				cache[c].register[registerIndex(a)] = 0
			}
			sort.Sort(cache)
			i, j := 0, 1
			for j < len(cache) {
				if cache[i].register == cache[j].register {
					cache[i].min = math.Min(cache[i].min, cache[j].min).(uint)
					cache[i].max = math.Max(cache[i].max, cache[j].max).(uint)
				} else {
					i += 1
					cache[i] = cache[j]
				}
				j += 1
			}
			if i < len(cache) {
				cache = cache[:i+1]
			}
			next := Cache{}
			for _, c := range cache {
				for i := 1; i <= 9; i++ {
					n := State{
						register: c.register,
						min:      c.min*10 + uint(i),
						max:      c.max*10 + uint(i),
					}
					n.register[registerIndex(a)] = i
					next = append(next, n)
				}
			}
			cache = next
			fmt.Printf("\t%d\n", len(cache))
		case "add":
			for c := range cache {
				cache[c].register[registerIndex(a)] += registerValue(b, cache[c].register)
			}
		case "mul":
			for c := range cache {
				cache[c].register[registerIndex(a)] *= registerValue(b, cache[c].register)
			}
		case "div":
			for c := range cache {
				cache[c].register[registerIndex(a)] /= registerValue(b, cache[c].register)
			}
		case "mod":
			for c := range cache {
				cache[c].register[registerIndex(a)] %= registerValue(b, cache[c].register)
			}
		case "eql":
			for c := range cache {
				A, B := cache[c].register[registerIndex(a)], registerValue(b, cache[c].register)
				if A == B {
					cache[c].register[registerIndex(a)] = 1
				} else {
					cache[c].register[registerIndex(a)] = 0
				}
			}
		}
	}
	min, max = 99999999999999, 0
	for _, c := range cache {
		if c.register[3] == 0 {
			max = math.Max(c.max, max).(uint)
			min = math.Min(c.min, min).(uint)
		}
	}
	return max, min
}
