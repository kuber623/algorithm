package leetcode_261

// https://leetcode.cn/problems/graph-valid-tree
// 难度：中等
// 题解：
// 使用并查集判断所有节点的连通性，当连通分量大于 1 时不为树
// 同时使用并查集判断添加一条边是否会形成环

func validTree(n int, edges [][]int) bool {
	set := NewUnionSet(n)

	for _, edge := range edges {
		// 判断添加一条边是否会形成环
		if set.Find(edge[0]) == set.Find(edge[1]) {
			return false
		}

		set.Union(edge[0], edge[1])
	}

	// 判断连通分量是否大于 1
	root := set.Find(0)
	for i := 1; i < n; i++ {
		if set.Find(i) != root {
			return false
		}
	}

	return true
}

type UnionSet struct {
	parents []int
	ranks   []int
}

func NewUnionSet(n int) *UnionSet {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}

	ranks := make([]int, n)
	for i := range ranks {
		ranks[i] = 1
	}

	return &UnionSet{parents, ranks}
}

func (set *UnionSet) Find(x int) int {
	px := set.parents[x]
	if px == x {
		return x
	}

	root := set.Find(px)
	set.parents[x] = root
	return root
}

func (set *UnionSet) Union(x, y int) {
	rx := set.Find(x)
	ry := set.Find(y)

	if rx == ry {
		return
	}

	if set.ranks[rx] < set.ranks[ry] {
		set.parents[rx] = ry
	} else if set.ranks[rx] > set.ranks[ry] {
		set.parents[ry] = rx
	} else {
		set.parents[ry] = rx
		set.ranks[rx]++
	}
}
