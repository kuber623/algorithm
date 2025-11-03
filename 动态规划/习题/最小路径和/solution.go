package leetcode_64

// https://leetcode.cn/problems/minimum-path-sum
// 难度：中等

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// dp[i][j] 表示到达 (i, j) 的最小数字和
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = grid[i][j]
			case i != 0 && j == 0:
				dp[i][j] = dp[i-1][j] + grid[i][j]
			case i == 0 && j != 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
			default:
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}
