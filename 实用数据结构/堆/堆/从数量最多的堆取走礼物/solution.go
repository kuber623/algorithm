package leetcode_2558

import (
	"container/heap"
	"math"
	"sort"
)

// https://leetcode.cn/problems/take-gifts-from-the-richest-pile/
// 难度：简单
// 题解：
// 使用最大堆实现快速挑选礼物堆，然后对礼物数量求平方根后重新进堆，执行 K 次
// 当堆的最大值为 1 时中断执行

func pickGifts(gifts []int, k int) int64 {
	h := &MaxHeap{gifts}
	heap.Init(h) // 原地堆化

	for ; k > 0; k-- {
		gifts[0] = int(math.Sqrt(float64(gifts[0]))) // 直接修改堆顶
		heap.Fix(h, 0)
	}

	ans := int64(0)
	for _, gift := range gifts {
		ans += int64(gift)
	}
	return ans
}

type MaxHeap struct {
	sort.IntSlice
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *MaxHeap) Push(x any) {}

func (h *MaxHeap) Pop() any {
	return nil
}
