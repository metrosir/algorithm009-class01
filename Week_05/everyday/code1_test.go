package everyday_test

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	if nums[0]+nums[1]+nums[2]+nums[3] > target || nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3]+nums[len(nums)-4] < target {
		return res
	}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] && i < len(nums)-3 {
			continue
		}
		for j := i + 2; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] && j < len(nums)-2 {
				continue
			}
			for start, end := j+1, len(nums)-1; start < end; {
				if start > j+1 && nums[start] == nums[start-1] && start < end {
					start++
					continue
				}
				if end < len(nums)-1 && nums[end] == nums[end+1] && start < end {
					end--
					continue
				}
				sum := nums[start] + nums[end] + nums[i] + nums[j]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
				}
				if sum < target {
					start++
				} else {
					end--
				}
			}

		}
	}
	return res
}

func TestFourS(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(tt(nums, target))
}

func tt(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	if len(nums) < 4 {
		return result
	}
	min := nums[0] + nums[1] + nums[2] + nums[3]
	max := nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3] + nums[len(nums)-4]
	if min > target || max < target {
		return result
	}
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] && i < len(nums)-3 {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] && j < len(nums)-2 {
				continue
			}
			for start, end := j+1, len(nums)-1; start < end; {
				if start > j+1 && nums[start] == nums[start-1] && start < end {
					start++
					continue
				}
				if end < len(nums)-1 && nums[end] == nums[end+1] && start < end {
					end--
					continue
				}

				sum := nums[start] + nums[end] + nums[i] + nums[j]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
				}
				if sum < target {
					start++
				} else {
					end--
				}
			}
		}
	}
	return result
}
