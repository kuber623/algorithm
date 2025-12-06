package leetcode_767

import "container/heap"

// https://leetcode.cn/problems/reorganize-string/
// 难度：中等

func reorganizeString(s string) string {
	ans := make([]rune, 0, len(s))

	counter := make(map[rune]int)
	for _, ch := range s {
		counter[ch]++
	}

	queue := &priorityQueue{
		items: make([]*item, 0, 26),
	}
	for ch, cnt := range counter {
		heap.Push(queue, &item{ch, cnt})
	}

	last := new(item)
	for queue.Len() > 0 {
		x := heap.Pop(queue).(*item)
		x.cnt--

		ans = append(ans, x.ch)
		if last != nil && last.cnt > 0 {
			heap.Push(queue, last)
		}
		last = x
	}

	if len(ans) < len(s) {
		return ""
	}
	return string(ans)
}

type priorityQueue struct {
	items []*item
}

func (q *priorityQueue) Len() int {
	return len(q.items)
}

func (q *priorityQueue) Less(i, j int) bool {
	return q.items[i].cnt > q.items[j].cnt
}

func (q *priorityQueue) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *priorityQueue) Push(x any) {
	q.items = append(q.items, x.(*item))
}

func (q *priorityQueue) Pop() any {
	o := q.items
	x := q.items[len(o)-1]
	q.items = o[:len(o)-1]
	return x
}

type item struct {
	ch  rune
	cnt int
}
