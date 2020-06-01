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
	fmt.Println(getLeastNumberss([]int{3, 2, 1}, 2))
}

//4、5、1、6、2、7、3、8
//4
//i+1 : r
//i-1 :
func getLeastNumberss(arr []int, k int) []int {
	if len(arr) < k {
		return nil
	}

	res := []int{}
	for i := 0; i < k; i++ {
		min := arr[0]
		for i := 1; i < len(arr); i++ {
			if min > arr[i] {
				min = arr[i]
			}
		}
		if i-1 < 0 {
			//arr = arr[i+2:]
			fmt.Println(1, arr[i+2:])
		} else if i+1 == len(arr) {
			//arr = arr[:len(arr)-1]
			fmt.Println(2, arr[:len(arr)-1])
		} else {
			//arr = append(arr[:i-1], arr[i+1:]...)
			fmt.Println(3, arr[:i-1], arr[i+1:])
		}

		res = append(res, min)
	}
	return res
}
