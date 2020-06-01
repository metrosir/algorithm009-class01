package everyday_test

import (
	"testing"
)

//https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
//给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。
//
//对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。

func TestKids(t *testing.T) {
	//var a int
	//a = 11
	//str := "a"+ strconv.Itoa(a)
	//fmt.Println(str);
	//return
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	println(kidsWithCandies(candies, extraCandies))
}

//【解法1】暴力
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	//先求出最大值
	for i := 0; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	res := []bool{}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
