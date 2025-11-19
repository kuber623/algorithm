package leetcode_53

// https://leetcode.cn/problems/maximum-subarray/
// 难度：中等
// 题解：
// 我们用 dp[i] 代表以第 i 个数结尾的「连续子数组的最大和」，因此可以得到 dp[i] 的状态转移公式为：
//     dp[i] = max{dp[i−1] + nums[i], nums[i]}

func maxSubArray(nums []int) int {
	// 边界条件
	if len(nums) == 0 {
		return 0
	}

	pre, maxsum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		maxsum = max(maxsum, pre)
	}

	return maxsum
}
