package Week_02_test

import (
    "fmt"
    "testing"
)

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
//给定一个二叉树，返回它的中序 遍历。
//
//示例:
//
//输入: [1,null,2,3]

//      1
//       \
//       2
//      /
//     3
//
//输出: [1,3,2]



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func inorderTraversal(root *TreeNode) []int {
    res = make([]int, 0)
    help(root)
    return res
}

var res []int
func help(root *TreeNode)  {
    if root == nil {
        return
    }

    help(root.Left)
    res = append(res, root.Val)
    help(root.Right)
}



func TestIno(t *testing.T)  {
    tree_node := &TreeNode{
        Val:1,
        Right:&TreeNode{
            Val:2,
            Left:&TreeNode{
                Val:3,
            },
        },
    }
    fmt.Println(inorderTraversal(tree_node))
}
