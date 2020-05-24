package Week_01_test

import (
	"fmt"
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func TestTreeNaryNode(t *testing.T) {

}

func levelOrder(root *Node) [][]int {
	ret := [][]int{}

	if root == nil {
		return ret
	}

	x := []*Node{root}

	for i := 0; len(x) > 0; i++ {
		ret = append(ret, []int{})
		cn := len(x)
		fmt.Println(cn)
		//y := []*Node{}
		for j := 0; j < cn; j++ {
			//node_ := x[j]
			//ret[i] = append(ret[i], node_.Val)
			if x[j] != nil {
				ret[i] = append(ret[i], x[j].Val)
				for _, v := range x[j].Children {
					x = append(x, v)
				}
			}
		}
		x = x[cn:]
	}
	return ret
}
