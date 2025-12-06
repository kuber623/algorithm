package leetcode_2680

// https://leetcode.cn/problems/maximum-or/
// 难度：中等

func maximumOr(nums []int, k int) (ans int64) {
	n := len(nums)

	// 前缀与数组
	pre := make([]int, k+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] | nums[i]
	}

	// 后缀与数组
	suf := make([]int, k+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] | nums[i]
	}

	for i := 0; i < n; i++ {
		ans = max(ans, int64(pre[i]|suf[i+1]|(nums[i]<<k)))
	}

	return
}
