package leetcode_1186

// https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion
// 难度：中等

func maximumSum(arr []int) int {
	dp := [2]int{}
	dp[0], dp[1] = arr[0], 0
	maxsum := arr[0]
	for i := 1; i < len(arr); i++ {
		dp[0], dp[1] = max(dp[0], 0)+arr[i], max(dp[1]+arr[i], dp[0])
		maxsum = max(maxsum, dp[0], dp[1])
	}
	return maxsum
}
