package leetcode_1642

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/furthest-building-you-can-reach/
// 难度：中等

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &MaxHeap{}
	for i := 1; i < len(heights); i++ {
		gap := max(heights[i]-heights[i-1], 0)
		if gap == 0 {
			continue
		}
		// 尝试使用砖块
		heap.Push(h, gap)
		bricks -= gap
		// 如果砖块不够则反悔
		if bricks < 0 {
			// 对消耗砖块最多的楼层改用梯子
			bricks += heap.Pop(h).(int)
			ladders--
			// 如果梯子不够则无法移动
			if ladders < 0 {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}

type MaxHeap struct {
	sort.IntSlice
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *MaxHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *MaxHeap) Pop() any {
	o := h.IntSlice
	x := o[len(o)-1]
	h.IntSlice = o[:len(h.IntSlice)-1]
	return x
}
