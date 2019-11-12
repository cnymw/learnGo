package p64

import "math"

func minPathSum(grid [][]int) int {
	jLen, iLen := len(grid), len(grid[0])
	dp := make([][]int, jLen)
	for i := range dp {
		dp[i] = make([]int, iLen)
	}

	for i := 0; i < iLen; i++ {
		for j := 0; j < jLen; j++ {
			if i == 0 && j == 0 {
				dp[j][i] = grid[j][i]
			} else if j == 0 && i > 0 {
				dp[j][i] = dp[j][i-1] + grid[j][i]
			} else if i == 0 && j > 0 {
				dp[j][i] = dp[j-1][i] + grid[j][i]
			} else {
				dp[j][i] = int(math.Min(float64(dp[j-1][i]), float64(dp[j][i-1]))) + grid[j][i]
			}
		}
	}
	return dp[jLen-1][iLen-1]
}
