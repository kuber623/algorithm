package leetcode_128

// https://leetcode.cn/problems/longest-consecutive-sequence/
// 难度：中等
// 题解：
// 将数组存储到哈希表中，然后遍历数组中的数字 num 是否存在 num+1，如果存在则创建一条 num -> num+1 的边，记录在并查集当中
// 最后遍历并查集查看各个集合的大小，取出最大值

func longestConsecutive(nums []int) int {
	kth := 0
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num] = kth
		kth++
	}

	set := NewUnionSet(len(nums))
	// 相比于遍历原数组，遍历哈希表避免
	for num, _ := range m {
		if _, ok := m[num+1]; !ok {
			continue
		}
		set.Union(m[num], m[num+1])
	}

	counter := make(map[int]int, 0)
	for _, kth = range m {
		counter[set.Find(kth)]++
	}

	longest := 0
	for _, cnt := range counter {
		if cnt > longest {
			longest = cnt
		}
	}

	return longest
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
