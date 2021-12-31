package day24

import (
	"fmt"
	"sort"

	"github.com/jrhorner1/AoC/pkg/math"
)

func Hardcoded() (min, max uint) {
	fmt.Println("Permutations:")
	cache := Cache{{[4]int{0, 0, 0, 0}, 0, 0}}
	for i := 0; i < 14; i++ {
		next := Cache{}
		for _, c := range cache {
			for j := 1; j <= 9; j++ {
				n := State{
					register: c.register,
					min:      c.min*10 + uint(j),
					max:      c.max*10 + uint(j),
				}
				n.register[0] = j
				next = append(next, n)
			}
		}
		cache = next
		for c := range cache {
			w, x, y, z := &cache[c].register[0], &cache[c].register[1], &cache[c].register[2], &cache[c].register[3]
			block := [14][3]int{{1, 14, 14}, {1, 14, 2}, {1, 14, 1}, {1, 12, 13}, {1, 15, 5}, {26, -12, 5}, {26, -12, 5}, {1, 12, 9}, {26, -7, 3}, {1, 13, 13}, {26, -8, 2}, {26, -5, 1}, {26, -10, 11}, {26, -7, 8}}
			*x = (0 + *z) % 26
			*z /= block[i][0]
			if *x+block[i][1] == *w {
				*x = 0
			} else {
				*x = 1
			}
			*y = ((0 + 25) * *x) + 1
			*z *= *y
			*y = ((0 + *w) + block[i][2]) * *x
			*z += *y
		}
		for c := range cache {
			cache[c].register[0] = 0
		}
		sort.Sort(cache)
		a, b := 0, 1
		for b < len(cache) {
			if cache[a].register == cache[b].register {
				cache[a].min = math.Min(cache[a].min, cache[b].min).(uint)
				cache[a].max = math.Max(cache[a].max, cache[b].max).(uint)
			} else {
				a += 1
				cache[a] = cache[b]
			}
			b += 1
		}
		if a < len(cache) {
			cache = cache[:a+1]
		}
		fmt.Printf("\t%d\n", len(cache))
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
