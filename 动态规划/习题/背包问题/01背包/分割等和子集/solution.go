package leetcode_416

// https://leetcode.cn/problems/partition-equal-subset-sum/
// 难度：中等
// 题解：转换成 0/1 背包问题即可，背包大小为数组和除以2，如果数组和为奇数直接返回 false

func canPartition(nums []int) bool {
	n := len(nums)

	// 计算背包大小
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 0 {
		return false
	}
	target := sum / 2

	// dp[i][j] 表示是否存在 [0, i] 区间的物品凑出价值为 j 的方案
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true // 表示所有物品都不选的方案
	}
	dp[0][nums[0]] = true

	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
			}
		}
	}

	return dp[n-1][target]
}
