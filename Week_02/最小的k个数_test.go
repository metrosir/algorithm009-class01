package Week_02_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
//示例 1：
//
//输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]

func getLeastNumbers(arr []int, k int) []int {

	if len(arr) == k {
		return arr
	}

	for i := len(arr) / 2; i >= 0; i-- {
		heapify(i, len(arr), arr)
	}
	fmt.Println(arr)
	for i := len(arr) - 1; i >= len(arr)-k; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		heapify(0, i, arr)
	}
	fmt.Println(arr)
	return arr[len(arr)-k:]
}

func heapify(root int, length int, arr []int) {
	for root < length {
		min, left, right := root, root*2+1, root*2+2
		if left < length && arr[left] < arr[min] {
			min = left
		}
		if right < length && arr[right] < arr[min] {
			min = right
		}
		if min != root {
			arr[min], arr[root] = arr[root], arr[min]
			root = min
		} else {
			break
		}
	}
}

func TestGL(t *testing.T) {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}
