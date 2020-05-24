package Week_01_test

import (
	"fmt"
	"testing"
)

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func TestBi(t *testing.T) {
	tree := &TreeNode2{
		Val: 3,
		Left: &TreeNode2{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode2{
			Val: 20,
			Left: &TreeNode2{
				Val: 15,
			},
			Right: &TreeNode2{
				Val: 7,
			},
		},
	}
	ret := levelOrder2(tree)
	fmt.Println(ret)
}

func levelOrder2(root *TreeNode2) [][]int {

	ret := [][]int{}
	tree := []*TreeNode2{root}
	for i := 0; i < len(tree); i++ {
		//node := []*TreeNode{}
		ret = append(ret, []int{})
		tree_ := []*TreeNode2{}
		for j := 0; j < len(tree); j++ {
			node := tree[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				tree_ = append(tree_, node.Left)
			}

			if node.Right != nil {
				tree_ = append(tree_, node.Right)
			}
		}
		tree = tree_
	}
	return ret
}
