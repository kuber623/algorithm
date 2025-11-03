package leetcode_97

// https://leetcode.cn/problems/interleaving-string
// 难度：中等

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			k := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (s1[i-1] == s3[k] && dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (s2[j-1] == s3[k] && dp[i][j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

// 由于数组 dp 的第 i 行只和第 i−1 行相关，因此可以使用滚动数组优化空间复杂度。
func isInterleaveOptimize(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([]bool, len(s2)+1)
	dp[0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			k := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[k]
			}
			if j > 0 {
				dp[j] = dp[j] || (s2[j-1] == s3[k] && dp[j-1])
			}
		}
	}

	return dp[len(s2)]
}
