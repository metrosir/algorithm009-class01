package everyday_test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]

//【解法1】双层循环 O(n2)
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	head_ := []ListNode{*head}
	res := []int{}
	for len(head_) > 0 {
		for _, v := range head_ {
			res = append([]int{v.Val}, res...)
			if v.Next != nil {
				head_ = []ListNode{*v.Next}
			} else {
				head_ = []ListNode{}
			}
		}
	}
	return res
}

func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := []int{}
	for head != nil {
		res = append([]int{head.Val}, res...)
		head = head.Next
	}
	return res
}

func new() *ListNode {

	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
}

func TestName(t *testing.T) {
	list := new()
	fmt.Println(reversePrint2(list))
}
