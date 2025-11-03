package union_set

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

func (set *UnionSet) Union(x, y int) {
	rx := set.Find(x)
	ry := set.Find(y)
	if rx == ry {
		return
	}

	// 按秩合并
	if set.ranks[rx] > set.ranks[ry] {
		set.parents[ry] = rx
	} else if set.ranks[rx] < set.ranks[ry] {
		set.parents[rx] = ry
	} else {
		set.parents[rx] = ry
		set.ranks[ry]++
	}
}

func (set *UnionSet) Find(x int) int {
	px := set.parents[x]
	if x == px {
		return x
	}

	// 路径压缩
	root := set.Find(px)
	set.parents[x] = root
	return root
}
