package leetcode_3418

// https://leetcode.cn/problems/maximum-amount-of-money-robot-can-earn/
// 难度：中等

func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])

	// dp[i][j][k] 表示机器人到达 (i, j) 后剩余 k 次感化能力的所获得的最大金币数
	dp := make([][][3]int, m)
	for i := range dp {
		dp[i] = make([][3]int, n)
		for j := range dp[i] {
			dp[i][j][0], dp[i][j][1], dp[i][j][2] = -100007, -100007, -100007
		}
	}
	dp[0][0][2], dp[0][0][1], dp[0][0][0] = coins[0][0], max(coins[0][0], 0), -100007

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cost := coins[i][j]

			if i > 0 {
				dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][0]+cost, dp[i-1][j][1])
				dp[i][j][1] = max(dp[i][j][1], dp[i-1][j][1]+cost, dp[i-1][j][2])
				dp[i][j][2] = max(dp[i][j][2], dp[i-1][j][2]+cost)
			}

			if j > 0 {
				dp[i][j][0] = max(dp[i][j][0], dp[i][j-1][0]+cost, dp[i][j-1][1])
				dp[i][j][1] = max(dp[i][j][1], dp[i][j-1][1]+cost, dp[i][j-1][2])
				dp[i][j][2] = max(dp[i][j][2], dp[i][j-1][2]+cost)
			}
		}
	}

	return max(dp[m-1][n-1][0], dp[m-1][n-1][1], dp[m-1][n-1][2])
}
