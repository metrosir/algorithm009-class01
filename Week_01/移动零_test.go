package Week_01_test

import (
	"fmt"
	"testing"
)

//link: https://leetcode-cn.com/problems/move-zeroes/
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

//【解法1】O(n2) O(1)
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

//【解法2】O(n) O(1)
func moveZeroesV2(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := 0; i < len(nums)-j; i++ {
		nums[j+i] = 0
	}
	return nums
}

//【解法3】O(n) O(1)
func moveZeroesV3(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			j++
		}
	}
	return nums
}

func TestZer(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroesV3(nums))
}
