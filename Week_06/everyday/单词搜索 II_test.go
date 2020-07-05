package everyday_test

import (
	"fmt"
	"testing"
)

type Node struct {
	word    string
	childen [26]*Node
}

func findWords(board [][]byte, words []string) []string {
	root := &Node{}
	for _, w := range words {
		node := root
		for _, c := range w {
			fmt.Println(c, 'a', c-'a')
			if node.childen[c-'a'] == nil {
				node.childen[c-'a'] = &Node{}
			}
			node = node.childen[c-'a']
		}
		node.word = w
	}
	res := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, root, &res)
		}
	}
	return res
}

func dfs(i, j int, board [][]byte, node *Node, res *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.childen[c-'a'] == nil {
		return
	}
	node = node.childen[c-'a']

	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}
	board[i][j] = '#'
	dfs(i+1, j, board, node, res)
	dfs(i-1, j, board, node, res)
	dfs(i, j+1, board, node, res)
	dfs(i, j-1, board, node, res)
	board[i][j] = c
}

func TestFind(t *testing.T) {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	fmt.Println(findWords(board, []string{"oath", "pea", "eat", "rain"}))
}
