package Week_03_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/subsets/

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//

//【解法1】 递归 回溯
//【解法2】 迭代

func subsets(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, 0, []int{}, &res)
	return res
}
func dfs(nums []int, start int, sol []int, res *[][]int) {
	n := len(nums)
	if start <= n {
		tmp := make([]int, len(sol))
		copy(tmp, sol)
		*res = append(*res, tmp)
	}
	for i := start; i < n; i++ {
		sol = append(sol, nums[i])
		dfs(nums, i+1, sol, res)
		sol = sol[:len(sol)-1]
	}
}

func dfs2(nums []int, start int, sol []int, res *[][]int) {
	tmp := make([]int, len(sol))
	copy(tmp, sol)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		sol = append(sol, nums[i])
		dfs2(nums, i+1, sol, res)
		sol = sol[:len(sol)-1]
	}
}

//回溯模板
/**
func dfs (step int) {
    if 判断边界条件 {
        相应逻辑
        return
    }

    for 尝试每种可能 {
        满足check条件
        标记
        下探到下一层 dfs(step + 1)
        恢复初始状态（回溯的时候要用到）
    }
}
*/

func subsetss(nums []int) [][]int {
	res := [][]int{}
	//opt := []int{}
	var dfs func(int, []int)
	dfs = func(idx int, sol []int) {
		if idx <= len(nums) {
			tmp := make([]int, len(sol))
			copy(tmp, sol)
			res = append(res, tmp)
		}
		for i := idx; i < len(nums); i++ {
			sol = append(sol, nums[i])
			dfs(i+1, sol)
			sol = sol[:len(sol)-1]
		}
	}
	dfs(0, []int{})
	return res
}

func TestSub(t *testing.T) {

	nums := []int{1, 2, 3}
	fmt.Println(subset(nums))
}

func subset(nums []int) [][]int {
	res := [][]int{}
	dfs3(&res, 0, []int{}, nums)
	return res
}

func dfs3(res *[][]int, idx int, sol []int, nums []int) {
	if idx <= len(nums) {
		tmp := make([]int, len(sol))
		copy(tmp, sol)
		*res = append(*res, tmp)
	}

	for i := idx; i < len(nums); i++ {
		sol = append(sol, nums[i])
		dfs3(res, i+1, sol, nums)
		sol = sol[:len(sol)-1]
	}
}
