package math

import (
	"math"
)

func getFibonacci(n int, cache *map[int]int) int {
	if n < 2 {
		(*cache)[n] = n
		return n
	}
	if _, found := (*cache)[n-1]; !found {
		(*cache)[n-1] = getFibonacci(n-1, cache)
	}
	if _, found := (*cache)[n-2]; !found {
		(*cache)[n-2] = getFibonacci(n-2, cache)
	}
	return (*cache)[n-1] + (*cache)[n-2]
}

func Fibonacci(n int) int {
	cache := make(map[int]int)
	result := make([]int, n)
	for i := 1; i < n; i++ {
		result[i-1] = getFibonacci(i, &cache)
	}
	return result[n-1]
}

func FibonacciSeq(n int) []int {
	cache := make(map[int]int)
	result := make([]int, n)
	for i := 1; i < n; i++ {
		result[i-1] = getFibonacci(i, &cache)
	}
	return result
}

func IntAbs(i int) int {
	return int(math.Abs(float64(i)))
}

func IntMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func IntMin(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
