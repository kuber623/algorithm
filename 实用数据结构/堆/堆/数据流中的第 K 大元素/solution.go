package leetcode_703

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/kth-largest-element-in-a-stream/
// 难度：简单

type KthLargest struct {
	minHeap *MinHeap
	kth     int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	h.IntSlice = make(sort.IntSlice, 0, k)

	for _, num := range nums {
		if len(h.IntSlice) < k {
			heap.Push(h, num)
			continue
		}

		if h.IntSlice[0] >= num {
			continue
		}

		h.IntSlice[0] = num
		heap.Fix(h, 0)
	}

	return KthLargest{minHeap: h, kth: k}
}

func (c *KthLargest) Add(val int) int {
	k, h := c.kth, c.minHeap

	if len(h.IntSlice) < k {
		heap.Push(h, val)
		return h.IntSlice[0]
	}

	if h.IntSlice[0] < val {
		h.IntSlice[0] = val
		heap.Fix(h, 0)
	}

	return h.IntSlice[0]
}

type MinHeap struct {
	sort.IntSlice
}

func (h *MinHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *MinHeap) Pop() any {
	return nil
}
