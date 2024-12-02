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

func Average(a *[]int64) int64 {
	sum := int64(0)
	for _, n := range *a {
		sum += n
	}
	return sum / int64(len(*a))
}

func LCM(a, b int64) int64 {
	gcd, lcm := int64(1), int64(1)
	if a > b {
		gcd = GCD(b, a)
		lcm = b * (a / gcd)
	} else if b > a {
		gcd = GCD(a, b)
		lcm = a * (b / gcd)
	}
	return lcm
}

func GCD(l, g int64) int64 {
	gcd := l
	r := g % l
	if r != 0 {
		gcd = GCD(r, l)
	}
	return gcd
}

func factorial(i int) int {
	if i > 1 {
		return i * factorial(i-1)
	}
	return 1
}

func combinations(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
