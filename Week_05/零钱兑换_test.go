package Week_05_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/coin-change/
//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//
//
//示例 1:
//
//输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1
//示例 2:
//
//输入: coins = [2], amount = 3
//输出: -1

//假设最小组合中最后一枚硬币是c：
//
//c 可能是 coins 中任何一个；
//去除 c 后剩下的部分，一定也是当前总额的一个最小组合（否则加上c不可能构成最小组合）
//或者用以下思路：
//如果 dp[i] 表示组成金额i的最小组合，dp[i]+1 一定是组成金额 i+c 的最小组合。

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//从1开始循环硬币
	for i := 1; i <= amount; i++ {

		//初始化当前硬币值
		dp[i] = -1

		//循环给定硬币的组合
		for _, c := range coins {

			//当前的硬币如果小于组合则跳过
			if i < c || dp[i-c] == -1 {
				continue
			}
			//如果组合中的一个数大于当前硬币，组合数前去当前硬币->剩余硬币的值 +1
			count := dp[i-c] + 1
			//如果当前硬币等于 -1。。等于初始值 ，或者当前硬币大于当前值
			fmt.Println(dp)
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}

//dp[1]=-1
//m1
//dp[1-1] = dp[0] = 0
//dp[1-1]+1 = 0+1
//dp[1] == -1 || dp[1] > 1 {
//    dp[1] = 1
//}
//dp [0 1 0 0]

//i=2
//2<1 || dp[2-1](1)
//count := dp[1](1) + 1
//if dp[2](0) == -1 || dp[2](0) > 2

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	//coins := []int{2}
	amount := 3

	fmt.Println(coinChange(coins, amount))
}
