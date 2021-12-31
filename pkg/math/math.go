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

func Abs(i interface{}) interface{} {
	switch i.(type) {
	case int:
		if i.(int) < 0 {
			return -i.(int)
		} else {
			return i.(int)
		}
	case int64:
		if i.(int64) < 0 {
			return -i.(int64)
		} else {
			return i.(int64)
		}
	case float64:
		if i.(float64) < 0 {
			return -i.(float64)
		} else {
			return i.(float64)
		}
	}
	return nil
}

func MaxInt() int {
	return int(math.MaxInt64)
}

func MaxUint() uint {
	return uint(math.MaxUint64)
}

func Max(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return a.(int)
		} else {
			return b.(int)
		}
	case uint:
		if a.(uint) > b.(uint) {
			return a.(uint)
		} else {
			return b.(uint)
		}
	}
	return nil
}

func Min(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		if a.(int) < b.(int) {
			return a.(int)
		} else {
			return b.(int)
		}
	case uint:
		if a.(uint) < b.(uint) {
			return a.(uint)
		} else {
			return b.(uint)
		}
	}
	return nil
}
