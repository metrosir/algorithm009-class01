package everyday_test

import (
	"fmt"
	"testing"
)

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	min, max := 0, num/2+1
	for max-min >= 0 {
		mid := (min + max) / 2
		if num > mid*mid {
			min = mid + 1
		} else if num < mid*mid {
			max = mid - 1
		} else if num == mid*mid {
			return true
		}
	}
	return false
}

func TestIsPerf(t *testing.T) {
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(2))
	fmt.Println(isPerfectSquare(4))
	fmt.Println(isPerfectSquare(14))
}
