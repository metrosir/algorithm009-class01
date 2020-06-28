package Week_06_test

import (
	"fmt"
	"testing"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	len_v := len(matrix[0]) - 1
	for line := 0; line < len(matrix); line++ {
		if matrix[line][len_v] >= target {
			for i := 0; i <= len_v; i++ {
				if matrix[line][i] == target {
					return true
				}
			}
		}
	}
	return false
}

func TestSearch(t *testing.T) {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	target := 3
	fmt.Println(searchMatrix(matrix, target))

}
