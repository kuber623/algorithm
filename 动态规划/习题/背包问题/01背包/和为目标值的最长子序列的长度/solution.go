package leetcode_2915

import "math"

// https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/
// 难度：中等
// 题解：0/1 背包问题 + 长度记录 + 滚动数组

func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)

	// dp[i][j] 表示是否存在 [0, i] 区间的物品凑出价值为 j 的方案
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true // 表示所有物品都不选的方案
	}
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	// length[i][j] 表示存在 [0, i] 区间的物品凑出价值为 j 的方案时的最大长度
	length := make([][]int, n)
	for i := range length {
		length[i] = make([]int, target+1)
	}
	if nums[0] <= target {
		length[0][nums[0]] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j < nums[i] {
				dp[i][j], length[i][j] = dp[i-1][j], length[i-1][j]
				continue
			}

			u, v := dp[i-1][j], dp[i-1][j-nums[i]]
			dp[i][j] = u || v
			switch {
			case u && !v:
				length[i][j] = length[i-1][j]
			case !u && v:
				length[i][j] = length[i-1][j-nums[i]] + 1
			case u && v:
				length[i][j] = max(length[i-1][j], length[i-1][j-nums[i]]+1)
			}
		}
	}

	if dp[n-1][target] {
		return length[n-1][target]
	}
	return -1
}

func lengthOfLongestSubsequenceOptimized(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = math.MinInt
	}

	s := 0
	for _, x := range nums {
		s = min(s+x, target)
		for j := s; j >= x; j-- {
			dp[j] = max(dp[j], dp[j-x]+1)
		}
	}

	if dp[target] > 0 {
		return dp[target]
	}
	return -1
}
