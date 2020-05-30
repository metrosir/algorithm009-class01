package Week_02_test

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
//【题意】
//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
//输入: [1,null,2,3]
//1
//\
//2
///
//3
//
//输出: [1,2,3]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//解法【1】
func preorderTraversal(root *TreeNode) []int {

	resV2 = make([]int, 0)
	helpV2(root)
	return resV2
}

var resV2 []int

func helpV2(root *TreeNode) {
	if root == nil {
		return
	}
	resV2 = append(resV2, root.Val)
	helpV2(root.Left)
	helpV2(root.Right)
}

func TestP(t *testing.T) {
	tree_node := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(preorderTraversal(tree_node))
}
