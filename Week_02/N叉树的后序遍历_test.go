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

 type Node struct {
     Val int
     Children []*Node
 }

func postorders(root *Node) []int {
    res := make([]int, 0)

    helps(root, &res)
    return res
}

func helps(root *Node, res *[]int)  {
    if root  == nil {
        return
    }

    for _, val := range root.Children{
        helps(val, res)
    }
    *res = append(*res, root.Val)
}

func postorder(root *Node) []int  {

    if root == nil {
        return nil
    }

    stack := make([]*Node, 0)
    stack = append(stack, root)
    res := make([]int, 0)
    for len(stack) > 0  {
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        for i:=0; i<len(node.Children) ; i++  {
            stack = append(stack,node.Children[i])
        }
    }
    return res
}

func TestPs(t *testing.T)  {
    v := &Node{
        Val:1,
        Children:[]*Node{
            &Node{Val:3,Children:[]*Node{&Node{Val:5},&Node{Val:6}}},&Node{Val:2},&Node{Val:4},},
    }
    fmt.Println(postorders(v))
}
