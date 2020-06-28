package Week_06_test

import (
	"fmt"
	"testing"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int
	m = len(obstacleGrid)
	if m > 0 {
		n = len(obstacleGrid[0])
	}

	dp := make([]int, n)
	dp[0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]

}

func TestUnin(t *testing.T) {
	n := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(n))
}
