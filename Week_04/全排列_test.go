package Week_04_test

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/permutations/
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]

func permute(nums []int) [][]int {
	res := [][]int{}
	option := make([]int, len(nums))
	validate := make([]int, len(nums))
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, option...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if validate[i] == 0 {
				option[idx] = nums[i]
				validate[i] = 1
				dfs(idx + 1)
				validate[i] = 0
			}
		}
	}
	dfs(0)
	return res
}

//解法1，dfs，回溯，使用额外数组validate 判断当前节点是否被访问过
func permutes(nums []int) [][]int {
	res := [][]int{}
	validate := make([]int, len(nums))
	option := make([]int, len(nums))

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, option...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if validate[i] == 0 {
				validate[i] = 1
				option[idx] = nums[i]
				dfs(idx + 1)
				validate[i] = 0
			}
		}
	}
	dfs(0)
	return res
}

//解法2，dfs，回溯，不使用额外数组，始终保证idx右边的是可选值
func permutess(nums []int) [][]int {
	res := [][]int{}
	option := make([]int, len(nums))

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, option...))
			return
		}
		for i := idx; i < len(nums); i++ {
			option[idx] = nums[i]
			nums[i], nums[idx] = nums[idx], nums[i]
			dfs(idx + 1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	dfs(0)
	return res
}

func permute2(nums []int) [][]int {
	res := [][]int{}
	opt := make([]int, len(nums))
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, opt...))
			return
		}

		for i := idx; i < len(nums); i++ {
			opt[idx] = nums[i]
			nums[idx], nums[i] = nums[i], nums[idx]
			dfs(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}
	dfs(0)
	return res
}

//使用额外的数组进行记录
func permute4(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, &res, nil, make([]bool, len(nums)))
	return res
}

func dfs(nums []int, ret *[][]int, path []int, used []bool) {
	if len(path) == len(nums) {
		*ret = append(*ret, path)
		return
	}

	next := make([]int, len(path))
	copy(next, path)
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		next = append(next, nums[i])
		dfs(nums, ret, next, used)
		used[i] = false
		next = next[:len(next)-1]
	}
}

func TestPermute(t *testing.T) {
	fmt.Println(permute4([]int{1, 2, 3}))
}

//交换
func permute5(nums []int) [][]int {
	res := [][]int{}
	dfsSwap(&res, nums, 0)
	return res
}

func dfsSwap(res *[][]int, nums []int, idx int) {
	if len(nums)-1 == idx {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	//    交换：
	//        阶乘：3*2*1
	//        idx=0:3个分支
	//            第一次交换，固定第一个元素：* 3
	//        idx=1:6个分支
	//            第二次交换，固定第二个元素，在第一次的基础上再分支：* 2
	//        idx=2:6个分支
	//            第3个元素已经固定了，所以可以省略这步：* 1
	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i]
		dfsSwap(res, nums, idx+1)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}
