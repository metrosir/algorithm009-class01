package Week_04_test

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/#/description
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7

func levelOrder(root *TreeNode) [][]int {
	var res = [][]int{}

	root_node := []*TreeNode{root}
	for i := 0; len(root_node) > 0; i++ {
		tmp_node := []*TreeNode{}
		res = append(res, []int{})
		for j := 0; j < len(root_node); j++ {
			res[i] = append(res[i], root_node[j].Val)
			if root_node[j].Left != nil {
				tmp_node = append(tmp_node, root_node[j].Left)
			}
			if root_node[j].Right != nil {
				tmp_node = append(tmp_node, root_node[j].Right)
			}
		}
		root_node = tmp_node
	}
	return res
}

func TestLevelOrder(t *testing.T) {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
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
	fmt.Println(levelOrder(node))
}
