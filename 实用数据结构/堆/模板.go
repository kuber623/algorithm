package heap

import "math"

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	return &Heap{data: []int{0}}
}

func (h *Heap) Heapify(data []int) {
	h.data = append(h.data, data...)
	// 从最后一个非叶子节点进行下沉
	for i := parent(len(h.data) - 1); i > 0; i-- {
		h.bubbleDown(i)
	}
}

func (h *Heap) Push(x int) {
	h.data = append(h.data, x)
	h.bubbleUp(len(h.data) - 1)
}

func (h *Heap) Pop() int {
	if h.Empty() {
		return math.MinInt
	}
	x := h.data[1]
	h.data[1] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.bubbleDown(1)
	return x
}

func (h *Heap) Empty() bool {
	return len(h.data) == 1
}

func (h *Heap) bubbleUp(i int) {
	for {
		p := parent(i)
		if p == i || h.data[i] >= h.data[p] {
			break
		}
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
	}
}

func (h *Heap) bubbleDown(i int) {
	for {
		mini := i
		l, r := left(i), right(i)

		if h.withinRange(l) && h.data[l] < h.data[mini] {
			mini = l
		}
		if h.withinRange(r) && h.data[r] < h.data[mini] {
			mini = r
		}

		if mini == i {
			break
		}

		h.data[i], h.data[mini] = h.data[mini], h.data[i]
		i = mini
	}
}

func (h *Heap) withinRange(i int) bool {
	return i < len(h.data) && i > 0
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return i * 2
}

func right(i int) int {
	return (i * 2) + 1
}
