package linked_forward_star

type edge struct {
	to   int
	next int
}

type LinkedForwardStar struct {
	cnt   int
	heads []int
	edges []edge
}

func NewLinkedForwardStar(n int) *LinkedForwardStar {
	return &LinkedForwardStar{}
}

func (star *LinkedForwardStar) Add(edges [][]int) {}
