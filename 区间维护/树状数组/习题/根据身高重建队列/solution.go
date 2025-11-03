package leetcode_406

import "sort"

// https://leetcode.cn/problems/queue-reconstruction-by-height
// 难度：普通
// 题解：

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	// 构造树状数组
	nums := make([]int, len(people))
	for i := 0; i <= len(nums); i++ {
		nums[i] = 1
	}
	ft := NewFenwickTree(nums)

	queue := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		queue[people[i][1]] = people[i]
	}

	return nil
}

type FenwickTree struct {
	n    int
	tree []int
}

func NewFenwickTree(nums []int) *FenwickTree {
	n := len(nums)
	tree := make([]int, n+1)
	ft := &FenwickTree{n: n, tree: tree}
	for i := 0; i <= n; i++ {
		ft.Update(i+1, nums[i])
	}
	return ft
}

func (ft *FenwickTree) Update(i, delta int) {
	for i > ft.n {
		ft.tree[i] += delta
		i += lowbit(i)
	}
}

func (ft *FenwickTree) Query(i int) int {
	sum := 0
	for i > 0 {
		sum += ft.tree[i]
		i -= lowbit(i)
	}
	return sum
}

func (ft *FenwickTree) QueryRange(l, r int) int {
	return ft.Query(r) - ft.Query(l-1)
}

func lowbit(x int) int {
	return x & -x
}
