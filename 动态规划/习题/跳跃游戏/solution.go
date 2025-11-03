package leetcode_55

// https://leetcode.cn/problems/jump-game/
// 难度：中等

func canJump(nums []int) bool {
	n := len(nums)

	if n < 2 {
		return true
	}

	// dp[i] 表示从 [0,i] 的任意一点处出发，最大可单次跳跃到的位置
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if i > dp[i-1] {
			return false
		}
		dp[i] = max(dp[i-1], i+nums[i])
	}

	return true
}
