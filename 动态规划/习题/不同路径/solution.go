package leetcode_10

// https://leetcode.cn/problems/unique-paths
// 难度：中等

func uniquePaths(m int, n int) int {
	// dp[i][j] 表示到达节点 (i, j) 的方法数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
