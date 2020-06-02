package Week_03_test

import (
	"fmt"
	"testing"
)

//  https://leetcode-cn.com/problems/generate-parentheses/
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例：
//
//输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]

//深度优先遍历
func generateParenthesis(n int) []string {

	var res []string
	help("", 0, 0, n, &res)
	return res
}

func help(current string, left int, right int, n int, res *[]string) {

	if left == n && right == n {
		*res = append(*res, current)
		return
	}

	if left < n {
		help(current+"(", left+1, right, n, res)
	}

	if right < n && right < left {
		help(current+")", left, right+1, n, res)
	}
}

func TestG(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
