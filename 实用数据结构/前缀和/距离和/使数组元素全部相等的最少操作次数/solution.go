package leetcode_2602

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/
// 难度：中等
// 题解：排序 + 二分搜索 + 前缀和

func minOperations(nums []int, queries []int) (ans []int64) {
	n := len(nums)

	// 计算前缀和 sum
	// 其中 sum[i] 表示数组在区间 [0, i) 的累加和
	sort.Ints(nums)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	for _, query := range queries {
		k := query
		i := search(nums, k)
		step := k*i - sum[i] + sum[n] - sum[i] - (n-i)*k
		ans = append(ans, int64(step))
	}

	return ans
}

// 查找有序数组 nums 中首个大于等于 target 的元素索引
// 如果数组中所有数均小于 target 则返回数组长度
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
