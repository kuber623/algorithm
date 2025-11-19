package leetcode_152

// https://leetcode.cn/problems/maximum-product-subarray/
// 难度：中等

func maxProduct(nums []int) int {
	// 边界条件
	if len(nums) == 1 {
		return nums[0]
	}

	// dp[i][0] 和 dp[i][1] 分别代表以 nums[i] 结尾的最大正数整数和最小负数乘积
	dp := make([][2]int, len(nums))
	dp[0][0], dp[0][1] = nums[0], nums[0]
	maxprod := nums[0]
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] > 0:
			dp[i][0], dp[i][1] = max(dp[i-1][0]*nums[i], nums[i]), dp[i-1][1]*nums[i]
		case nums[i] < 0:
			dp[i][0], dp[i][1] = dp[i-1][1]*nums[i], min(dp[i-1][0]*nums[i], nums[i])
		default:
			dp[i][0], dp[i][1] = 0, 0
		}
		maxprod = max(maxprod, dp[i][0], dp[i][1])
	}

	return maxprod
}
