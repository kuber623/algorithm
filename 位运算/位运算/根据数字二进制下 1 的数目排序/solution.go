package leetcode_1356

import (
	"cmp"
	"math/bits"
	"slices"
)

// https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/
// 难度：简单

func sortByBits(arr []int) []int {
	slices.SortFunc(arr, func(a, b int) int {
		return cmp.Or(bits.OnesCount(uint(a))-bits.OnesCount(uint(b)), a-b)
	})
	return arr
}
