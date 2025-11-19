package leetcode_322

import "math"

// https://leetcode.cn/problems/coin-change/
// 难度：中等

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
