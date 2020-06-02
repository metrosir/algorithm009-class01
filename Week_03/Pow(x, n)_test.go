package Week_03_test

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/powx-n/

//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//示例 1:
//
//输入: 2.00000, 10
//输出: 1024.00000
//示例 2:
//
//输入: 2.10000, 3
//输出: 9.26100

// 2 * 2 * 2 * 2

// 4 * 4

func myPow(x float64, n int) float64 {

	if n >= 0 {
		return helps(x, n)
	}
	return 1 / helps(x, n)
}

func helps(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := helps(x, n/2)

	if n%2 == 0 {
		return y * y
	}

	return y * y * x
}

func TestMyPow(t *testing.T) {
	fmt.Println(myPow(2.1, 3))
}
