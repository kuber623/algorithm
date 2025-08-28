package heap

type Heap struct {
	data       []int
	comparator func(int, int) bool // Less
}

func NewHeap(comparator func(int, int) bool) *Heap {
	return &Heap{
		comparator: comparator,
	}
}

// Heapify 自下而上的建堆
func (h *Heap) Heapify(data []int) {
	h.data = data
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		h.bubbleDown(i)
	}
}

func (h *Heap) Enqueue(i int) {
	h.data = append(h.data, i)
	h.bubbleUp(len(h.data) - 1)
}

func (h *Heap) Dequeue() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	val := h.data[0]
	h.swap(0, len(h.data)-1)
	h.data = h.data[:len(h.data)-1]
	h.bubbleDown(0)
	return val, true
}

func (h *Heap) bubbleUp(i int) {
	p := h.parent(i)
	for i > 0 && h.comparator(h.data[i], h.data[p]) {
		h.swap(i, p)
		i = p
	}
}

func (h *Heap) bubbleDown(i int) {
	for h.withinRange(i) {
		l, r := h.left(i), h.right(i)
		switch {
		// 如果没有子节点
		case !h.withinRange(l):
			break
		// 如果只有左子节点
		case !h.withinRange(r):
			if h.comparator(h.data[i], h.data[l]) {
				h.swap(i, l)
				i = l
			}
		// 如果均有左右子节点
		default:
			if !h.comparator(h.data[i], h.data[l]) && !h.comparator(h.data[i], h.data[r]) {
				break
			}
			if h.comparator(h.data[l], h.data[r]) {
				h.swap(i, l)
				i = l
			} else {
				h.swap(i, r)
				i = r
			}
		}
	}
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) left(i int) int {
	return 2*i + 1
}

func (h *Heap) right(i int) int {
	return 2*i + 2
}

func (h *Heap) withinRange(i int) bool {
	return i >= 0 && i < len(h.data)
}
