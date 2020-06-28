package Week_06_test

import (
	"fmt"
	"testing"
)

var cache = map[int]int{}

func climbStairs(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	cache[n] = climbStairs(n-1) + climbStairs(n-2) + climbStairs(n-3)
	return cache[n]
}

func TestClimbSt(t *testing.T) {
	fmt.Println(climbStairs(4))
}
