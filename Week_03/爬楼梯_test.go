package Week_03_test

import (
	"fmt"
	"testing"
)

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶

//1 1
//2 2
//3 111 12 21 3
//4 1111 121 112 211  22 5
//n = (n - 1) + (n-2)
var cache = map[int]int{}

func climbStairs(n int) int {
	if n <= 2 {
		cache[n] = n
		return n
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}

func TestClim(t *testing.T) {
	fmt.Println(climbStairs(5))
}
