package everyday

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/get-kth-magic-number-lcci/
//有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
//
//示例 1:
//
//输入: k = 5
//
//输出: 9
//

func getKthMagicNumber(k int) int {
	res := map[int]int{}
	p3 := 0
	p5 := 0
	p7 := 0
	res[0] = 1
	for i := 1; i <= k; i++ {
		res[i] = min([]int{res[p3] * 3, res[p5] * 5, res[p7] * 7})
		if res[i] == res[p3]*3 {
			p3++
		}
		if res[i] == res[p5]*5 {
			p5++
		}
		if res[i] == res[p7]*7 {
			p7++
		}
	}
	return res[k-1]
}

func min(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr[0]
}

func TestKthM(t *testing.T) {
	fmt.Println(getKthMagicNumber(100))
}
