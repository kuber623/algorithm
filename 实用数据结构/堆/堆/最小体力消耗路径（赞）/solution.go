package leetcode_1631

import "container/heap"

// https://leetcode.cn/problems/path-with-minimum-effort/
// 难度：中等

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])

	// 用小根堆维护边权重
	eh := &EdgeHeap{edges: make([][3]int, 0)}
	for i := range heights {
		for j := range heights[i] {
			// 添加下方节点的边
			if i > 0 {
				heap.Push(eh, [3]int{i*n + j, (i-1)*n + j, abs(heights[i][j] - heights[i-1][j])})
			}
			// 添加左侧节点的边
			if j > 0 {
				heap.Push(eh, [3]int{i*n + j, i*n + j - 1, abs(heights[i][j] - heights[i][j-1])})
			}
		}
	}

	// 使用并查集维护连通性
	us := NewUnionSet(m * n)
	for eh.Len() > 0 {
		w := eh.edges[0][2]
		for eh.Len() > 0 && eh.edges[0][2] == w {
			edge := heap.Pop(eh).([3]int)
			us.Union(edge[0], edge[1])
		}
		if us.Find(0) == us.Find(m*n-1) {
			return w
		}
	}

	return 0
}

type UnionSet struct {
	parent []int
}

func NewUnionSet(n int) *UnionSet {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionSet{parent}
}

func (u *UnionSet) Union(x int, y int) {
	px, py := u.Find(x), u.Find(y)
	if px == py {
		return
	}
	u.parent[py] = px
}

func (u *UnionSet) Find(x int) int {
	p := u.parent[x]
	if p == x {
		return x
	}
	root := u.Find(p)
	u.parent[x] = root
	return root
}

type EdgeHeap struct {
	edges [][3]int
}

func (h *EdgeHeap) Len() int {
	return len(h.edges)
}

func (h *EdgeHeap) Less(i, j int) bool {
	return h.edges[i][2] < h.edges[j][2]
}

func (h *EdgeHeap) Swap(i, j int) {
	h.edges[i], h.edges[j] = h.edges[j], h.edges[i]
}

func (h *EdgeHeap) Push(x any) {
	h.edges = append(h.edges, x.([3]int))
}

func (h *EdgeHeap) Pop() any {
	o := h.edges
	x := o[len(o)-1]
	h.edges = o[:len(o)-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
