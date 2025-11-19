package leetcode_3026

import "math"

// https://leetcode.cn/problems/maximum-good-subarray-sum/
// 难度：中等
// 题解：
// 枚举 + 前缀和，前缀和用于快速计算子数组和，枚举用于判断是否存在好子数组

func maximumSubarraySum(nums []int, k int) int64 {
	ans := math.MinInt
	sum := 0
	minSum := make(map[int]int, len(nums))
	for _, num := range nums {
		// 判断是否存在好子数组
		s, ok := minSum[num+k]
		if ok {
			ans = max(ans, sum-s+num)
		}
		s, ok = minSum[num-k]
		if ok {
			ans = max(ans, sum-s+num)
		}

		// 更新当前数字的最小前缀和
		s, ok = minSum[num]
		if ok {
			minSum[num] = min(s, sum)
		} else {
			minSum[num] = sum
		}
		sum += num
	}

	// 如果不存在好子数组则返回 0
	if ans == math.MinInt {
		ans = 0
	}

	return int64(ans)
}
