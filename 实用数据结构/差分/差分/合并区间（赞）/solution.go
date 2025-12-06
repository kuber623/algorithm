package leetcode_56

import (
	"slices"
)

// https://leetcode.cn/problems/merge-intervals/
// 难度：中等
// 题解：差分解法

// 差分解法
func merge(intervals [][]int) (ans [][]int) {
	diff := make([]int, 20001)

	for _, interval := range intervals {
		l, r := interval[0], interval[1]
		diff[2*l]++
		diff[2*r+1]--
	}

	sum, pre := 0, -1
	for i := 0; i < len(diff); i++ {
		if diff[i] == 0 {
			continue
		}

		sum += diff[i]
		if pre == -1 {
			pre = i
		} else if sum == 0 {
			ans = append(ans, []int{pre / 2, i / 2})
			pre = -1
		}
	}
	return
}

// 合并解法
func mergeIntervals(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(i, j []int) int {
		return i[0] - j[0]
	})
	for _, interval := range intervals {
		n := len(ans)
		if n > 0 && interval[0] < ans[n-1][1] { // 区间相交，可以合并
			ans[n-1][1] = max(interval[1], ans[n-1][1])
		} else { // 区间不相交
			ans = append(ans, interval)
		}
	}
	return
}
