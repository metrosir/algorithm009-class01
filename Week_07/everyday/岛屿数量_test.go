package everyday_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/number-of-islands/

//dfs 算法
func numIslands(grid [][]byte) int {
	res := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 {
			return
		}
		if x >= len(grid) || y >= len(grid[0]) {
			return
		}
		if grid[x][y] == '0' {
			return
		}

		grid[x][y] = '0'
		dfs(x, y-1)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x+1, y)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}

func TestNum(t *testing.T) {
	ns := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(ns))
}
