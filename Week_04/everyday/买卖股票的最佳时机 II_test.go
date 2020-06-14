package everyday_test

import (
	"fmt"
	"testing"
)

func maxProfit(prices []int) int {
	var res = 0

	for i := 0; i < len(prices); i++ {
		if i+1 < len(prices) && prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1,2,3,4,5}
	fmt.Println(maxProfit(prices))
}
