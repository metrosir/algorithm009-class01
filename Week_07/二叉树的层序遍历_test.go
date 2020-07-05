package Week_07_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/submissions/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if res == nil {
		return res
	}
	node := []*TreeNode{root}
	for i := 0; len(node) > 0; i++ {
		res = append(res, []int{})
		tmp := []*TreeNode{}
		for j := 0; j < len(node); j++ {
			res[i] = append(res[i], node[j].Val)
			if node[j].Left != nil {
				tmp = append(tmp, node[j].Left)
			}
			if node[j].Right != nil {
				tmp = append(tmp, node[j].Right)
			}
		}
		node = tmp
	}
	return res
}

func TestBi(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ret := levelOrder(tree)
	fmt.Println(ret)
}
