package leetcode_63

// https://leetcode.cn/problems/unique-paths-ii
// 难度：中等

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

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
