package leetcode_746

// https://leetcode.cn/problems/min-cost-climbing-stairs
// 难度：简单
// 题解：
// 假设数组 cost 的长度为 n，则 n 个阶梯分别对应下标 0 到 n−1，楼层顶部对应下标 n，问题等价于计算达到下标 n 的最小花费。可以通过动态规划求解。
// 由于可以选择下标 0 或 1 作为初始阶梯，因此有 dp[0] = dp[1] = 0。
// 由于第 i 个台阶只能从第 i - 1 或 i - 2 个台阶攀爬到，因此可以得到其状态转移方程为：
//     dp[i]=min(dp[i−1]+cost[i−1],dp[i−2]+cost[i−2])
// 注意到 dp[i] 仅与 dp[i-1] 和 dp[i-2] 相关
// 因此可以通过「滚动数组」将空间复杂度由于到 O(1)

// 常规版本
func minCostClimbingStairs(cost []int) int {
	// dp[i] 表示到达第 i 个台阶线上爬所需要的费用
	dp := make([]int, len(cost)+1)
	// 由于到达第 0 和第 1 个台阶无需支付费用，估从第二个台阶开始计算费用
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}

// 滚动数组空间优化版本
func minCostClimbingStairsSpaceOptimize(cost []int) int {
	dp := [2]int{}
	for i := 2; i <= len(cost); i++ {
		dp[0], dp[1] = dp[1], min(dp[0]+cost[i-2], dp[1]+cost[i-1])
	}
	return dp[1]
}
