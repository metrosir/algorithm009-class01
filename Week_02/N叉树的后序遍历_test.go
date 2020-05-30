package Week_02_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//【解题方式】
//1、递归
//2、栈

type Node struct {
	Val      int
	Children []*Node
}

//【方法1】递归
func postorders(root *Node) []int {

	res := []int{}
	helps(root, &res)
	return res
}

func helps(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, v := range root.Children {
		helps(v, res)
	}
	*res = append(*res, root.Val)
}

//【方法2】栈
// [1]
// [3,2,4]
// [5,6]

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	//初始化栈/入栈/把顶节点入栈
	stack := []*Node{root}
	res := []int{}
	for len(stack) > 0 {
		//获取当前节点
		node := stack[len(stack)-1]
		//出栈
		stack = stack[:len(stack)-1]
		//保存当前节点的值
		res = append([]int{node.Val}, res...)
		for i := 0; i < len(node.Children); i++ {
			//批量入栈
			stack = append(stack, node.Children[i])
		}
	}
	return res
}

func TestPs(t *testing.T) {
	v := &Node{
		Val: 1,
		Children: []*Node{
			&Node{Val: 3,
				Children: []*Node{
					&Node{Val: 5}, &Node{Val: 6}}},
			&Node{Val: 2},
			&Node{Val: 4},
		},
	}
	fmt.Println(postorder(v))
}
