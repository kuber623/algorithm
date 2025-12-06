package leetcode_2589

import "slices"

// https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/
// 难度：困难

func findMinimumTime(tasks [][]int) int {
	slices.SortFunc(tasks, func(i, j []int) int {
		return i[0] - j[0]
	})

	return 0
}
