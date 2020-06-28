package everyday_test

import (
	"fmt"
	"testing"
)

func jump(nums []int) int {
	j := len(nums) - 1
	res := 0
	for j > 0 {
		for i := 0; i <= j; i++ {
			if nums[i]+i >= j {
				res++
				j = i
				break
			}
		}
	}
	return res
}

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
