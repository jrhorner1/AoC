package day6

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, days int) uint {
	in := strings.Split(strings.TrimSpace(string(*input)), ",")
	fish := map[int]uint{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, i := range in {
		out, _ := strconv.Atoi(i)
		fish[out] += 1
	}
	total := uint(0)
	for day := 0; day < days; day++ {
		tmp := fish[0]
		for i := 0; i < 8; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += tmp
		fish[8] = tmp
	}
	for _, v := range fish {
		total += v
	}
	return total
}
