package leetcode_3362

import (
	"container/heap"
	"slices"
	"sort"
)

// https://leetcode.cn/problems/zero-array-transformation-iii/
// 难度：中等
// 题解：
// 请看茶神题解 https://leetcode.cn/problems/zero-array-transformation-iii/solutions/2998650/tan-xin-zui-da-dui-chai-fen-shu-zu-pytho-35o6/

func maxRemoval(nums []int, queries [][]int) int {
	// 按照左端点从小到大进行排序
	slices.SortFunc(queries, func(a, b []int) int {
		return a[0] - b[0]
	})

	h, diff := hp{}, make([]int, len(nums)+1)
	sum, j := 0, 0
	for i, x := range nums {
		sum += diff[i]
		// 维护左端点小于 i 的区间
		for ; j < len(queries) && queries[j][0] <= i; j++ {
			heap.Push(&h, queries[j][1])
		}
		// 选择右端点最大的区间
		for sum < x && h.Len() > 0 && h.IntSlice[0] >= i {
			// 通过差分数组实现快速区间减一
			sum++
			diff[heap.Pop(&h).(int)+1]--
		}
		if sum < x {
			return -1
		}
	}

	return h.Len()
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
