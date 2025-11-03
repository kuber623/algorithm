package fenwick_tree

type FenwickTree struct {
	n    int   // 数组长度
	tree []int // 树状数组存储（索引从 1 开始方便二进制计算）
}

func NewFenwickTree(nums []int) *FenwickTree {
	n := len(nums)
	tree := make([]int, n+1)
	ft := &FenwickTree{n: n, tree: tree}
	for i := 0; i < n; i++ {
		ft.Update(i+1, nums[i])
	}
	return ft
}

// Update 单点更新：第 i 个元素增加 delta
func (ft *FenwickTree) Update(i, delta int) {
	for i <= ft.n {
		ft.tree[i] += delta
		i += lowbit(i)
	}
}

// Query 前缀和查询：前 i 个数之和
func (ft *FenwickTree) Query(i int) int {
	if i > ft.n || i < 0 {
		return 0
	}

	sum := 0
	for i > 0 {
		sum += ft.tree[i]
		i -= lowbit(i)
	}
	return sum
}

// QueryRange 区间和查询：区间 [L, R] 之和
func (ft *FenwickTree) QueryRange(L, R int) int {
	if L > R {
		return 0
	}
	return ft.Query(R) - ft.Query(L-1)
}

func lowbit(x int) int {
	return x & -x
}
