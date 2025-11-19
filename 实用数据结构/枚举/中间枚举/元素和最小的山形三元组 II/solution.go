package leetcode_2909

import "math"

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/
// 难度：中等
// 题解：枚举中间值，记录前缀最小值和后缀最小值，如果中间值比前缀最小值和后缀最小值都大，则满足山形三元组

func minimumSum(nums []int) int {
	ans := math.MaxInt

	n := len(nums)
	pref, suf := nums[0], make([]int, n)
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	for i := 1; i < n-1; i++ {
		// 判断是否满足
		if nums[i] > pref && nums[i] > suf[i+1] {
			ans = min(ans, nums[i]+pref+suf[i+1])
		}
		pref = min(pref, nums[i])
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
