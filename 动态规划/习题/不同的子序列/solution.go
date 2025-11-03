package leetcode_115

// https://leetcode.cn/problems/distinct-subsequences
// 难度：困难

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	// dp[i][j] 表示子串 s(0, i) 出现 t(O, j) 子序列的个数
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if s[i] == t[j] {
				switch {
				case i == 0 && j == 0:
					dp[i][j] = 1
				case i == 0:
					dp[i][j] = 0
				case j == 0:
					dp[i][j] = dp[i-1][j] + 1
				default:
					dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
				}
			} else {
				if i != 0 {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	return dp[m-1][n-1]
}
