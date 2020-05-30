package Week_02_test

import (
    "fmt"
    "testing"
)

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
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
    Val int
    Left *TreeNode
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
func preorderTraversal(root *TreeNode) []int {

    res = make([]int, 0)
    help(root)
    return res
}

var res []int
func help(root *TreeNode)  {
    if root == nil {
        return
    }
    res = append(res, root.Val)
    help(root.Left)
    help(root.Right)
}

func TestP(t *testing.T)  {
    tree_node := &TreeNode{
        Val:1,
        Right:&TreeNode{
            Val:2,
            Left:&TreeNode{
                Val:3,
            },
        },
    }
    fmt.Println(preorderTraversal(tree_node))
}


