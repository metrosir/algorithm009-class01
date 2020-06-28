package Week_04_test

import (
	"fmt"
	"testing"
)

func mySqrt(x int) int {

	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x
	mid := 1
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}

func TestMySqrt(t *testing.T) {

	fmt.Println(mySqrt(4))
}
