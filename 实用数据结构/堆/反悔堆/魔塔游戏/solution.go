package lcp_30

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/p0NxJO/
// 难度：中等

func magicTower(nums []int) (ans int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < 0 {
		return -1
	}

	hp, h := 1, &MinHeap{}
	for _, num := range nums {
		if num < 0 {
			heap.Push(h, num)
		}

		hp += num

		if hp < 1 {
			hp -= heap.Pop(h).(int)
			ans++
		}
	}
	return
}

type MinHeap struct{ sort.IntSlice }

func (h *MinHeap) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }

func (h *MinHeap) Pop() any { o := h.IntSlice; x := o[len(o)-1]; h.IntSlice = o[:len(o)-1]; return x }
