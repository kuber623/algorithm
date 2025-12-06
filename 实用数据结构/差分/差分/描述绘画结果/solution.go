package leetcode_1943

import (
	"sort"
)

// https://leetcode.cn/problems/describe-the-painting/
// 难度：中等
// 题解：使用哈希表和差分思想记录所有颜色变化

func splitPainting(segments [][]int) [][]int64 {
	ans := make([][]int64, 0)

	changes, axis := make(map[int]map[int]bool, 100000), make([]int, 0, 100000)
	for _, segment := range segments {
		begin, end, color := segment[0], segment[1], segment[2]

		if _, ok := changes[begin]; !ok {
			changes[begin] = make(map[int]bool)
			axis = append(axis, begin)
		}
		changes[begin][color] = true
		if _, ok := changes[end]; !ok {
			changes[end] = make(map[int]bool)
			axis = append(axis, end)
		}
		changes[end][-color] = true
	}
	sort.Ints(axis)

	pre, sum := -1, int64(0)
	for _, i := range axis {
		if sum > 0 {
			ans = append(ans, []int64{int64(pre), int64(i), sum})
		}

		for e := range changes[i] {
			sum += int64(e)
		}

		pre = i
	}

	return ans
}
