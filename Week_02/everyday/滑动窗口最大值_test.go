package Week_02_everyday_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/sliding-window-maximum/
//【题意】
//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。
//

//{1,3,-1,-3,5,3,6,7}

//len: 8

//range = len - ()
//2 = len - 0
//3 = len - 1
//4 = len - 2
//...
//8 = 1
// len - k + 1 层 / 个桶

//【解法1】暴力求解
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if len(nums) < k {
		return res
	}
	if k == 1 {
		return nums
	}
	t := len(nums) - k + 1
	for i := 0; i < t; i++ {
		max := 0
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

//【解法2】双端队列

func TestMaxSl(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
