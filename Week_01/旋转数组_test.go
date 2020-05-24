package Week_01_test

import (
	"fmt"
	"testing"
)

//link: https://leetcode-cn.com/problems/rotate-array/
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]

func rotate(nums []int, k int) {
	cn := len(nums)
	arr_tmp := []int{}
	for i := 0; i < k; i++ {
		arr_tmp = append(arr_tmp, nums[cn-1])
		cn--
	}
	for j := 0; j < len(nums)-k; j++ {
		arr_tmp = append(arr_tmp, nums[j])
	}
	nums = arr_tmp
	fmt.Println(nums)
}

func rotates(nums []int, k int) {
	for i := 0; i < k; i++ {
		r := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			tmp := nums[j]
			nums[j] = r
			r = tmp
		}
	}
	fmt.Println(nums)
}

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotates(arr, 3)
}
