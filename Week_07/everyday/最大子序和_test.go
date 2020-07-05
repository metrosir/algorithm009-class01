package everyday_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/maximum-subarray/

//动态规划
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

func TestMax(t *testing.T) {
	test_ := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(test_))
}
