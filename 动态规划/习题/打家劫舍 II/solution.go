package leetcode_213

// https://leetcode.cn/problems/house-robber-ii/
// 难度：中等
// 题解：
// 整体解法同打家劫舍一样，只不过由于是环形数组，故不能同时窃取数组中的第一个房间和最后一个房间
// 为此我们可以按照区间 [0, n−2] 和 [1, n−1] 进行两次计算，取其中的最大值得到最终结果
// 注意到 dp[i] 仅与 dp[i-1] 和 dp[i-2] 有关，因此可以使用滚动数组优化空间复杂度

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else {
		return max(doRob(nums[:len(nums)-1]), doRob(nums[1:]))
	}
}

func doRob(nums []int) int {
	// 滚动数组优化空间复杂度
	dp := [2]int{}
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[0], dp[1] = dp[1], max(dp[0]+nums[i], dp[1])
	}
	return dp[1]
}
