package leetcode_198

// https://leetcode.cn/problems/house-robber
// 难度：中等

func rob(nums []int) int {
	// 边界条件
	if len(nums) == 1 {
		return nums[0]
	}

	// dp[i] 表示区间 [0, i] 能窃取到的最高金额
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func robSpaceOptimize(nums []int) int {
	// 边界条件
	if len(nums) == 1 {
		return nums[0]
	}
	// 滚动数组优化空间复杂度
	dp := [2]int{}
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[0], dp[1] = dp[1], max(dp[0]+nums[i], dp[1])
	}
	return dp[1]
}
