package everyday_test

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/3sum/
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for start < end {
			tmp := nums[i] + nums[start] + nums[end]
			if tmp == 0 {
				res = append(res, []int{nums[i], nums[start], nums[end]})
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				start++
				end--
			} else if tmp < 0 {
				start++
			} else if tmp > 0 {
				end--
			}
		}
	}
	return res
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
