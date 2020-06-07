package Week_03_test

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/n-queens/

type queen struct {
	n     int
	stack map[int]int
	col   []int
	pie   map[int]int
	na    map[int]int
	ret   [][]string
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	if n == 0 {
		return res
	}
	q := NewQueen(n)
	q.backtrack(0)
	return q.ret
}

func NewQueen(n int) *queen {
	stack := make(map[int]int, n)
	col := make([]int, n)
	pie := make(map[int]int, n)
	na := make(map[int]int, n)
	return &queen{n, stack, col, pie, na, [][]string{}}
}

func (q *queen) backtrack(row int) {
	if q.n < 1 {
		return
	}
	for col := 0; col < q.n; col++ {
		if q.isNotUnderAttack(row, col) {
			q.addPos(row, col)
			if row+1 == q.n {
				q.addSolution()
			} else {
				q.backtrack(row + 1)
			}
			q.rmPos(row, col)
		}
	}
}

func (q *queen) isNotUnderAttack(row, col int) bool {
	if q.col[col]+q.pie[row+col]+q.na[col-row+2*q.n] == 0 {
		return true
	}
	return false
}

func (q *queen) addPos(row, col int) {
	q.stack[row] = col
	q.col[col] = 1
	q.pie[col+row] = 1
	q.na[col-row+2*q.n] = 1
}

func (q *queen) rmPos(row, col int) {
	delete(q.stack, row)
	q.col[col] = 0
	q.pie[row+col] = 0
	q.na[col-row+2*q.n] = 0
}

func (q *queen) addSolution() {
	solution := make([]string, q.n)
	for row := 0; row < q.n; row++ {
		rowCol := q.stack[row]
		output := ""
		for col := 0; col < q.n; col++ {
			if rowCol == col {
				output += "Q"
			} else {
				output += "."
			}
		}
		solution[row] = output
	}
	q.ret = append(q.ret, solution)
}

func TestSolve(t *testing.T) {
	fmt.Println(solveNQueens(5))
}
