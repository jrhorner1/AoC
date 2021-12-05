package math

import (
	"math"
)

func Fibonacci(target int) []uint {
	seq := []uint{0, 1}
	limit := 2
	for i := 1; i <= limit; i++ {
		next := seq[i] + seq[i-1]
		seq = append(seq, next)
		if next < uint(target) {
			limit++
		}
	}
	return seq
}

func IntAbs(i int) int {
	return int(math.Abs(float64(i)))
}
