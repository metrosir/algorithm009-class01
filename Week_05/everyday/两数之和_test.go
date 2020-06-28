package everyday_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/two-sum/description/
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

func twoSum(nums []int, target int) []int {
	tmp_map := map[int]int{}
	for i := 0; i < len(nums); i++ {
		x := target - nums[i]
		if _, ok := tmp_map[x]; !ok {
			tmp_map[nums[i]] = i
		} else {
			return []int{tmp_map[x], i}
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
