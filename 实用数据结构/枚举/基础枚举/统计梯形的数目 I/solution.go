package leetcode_3632

// https://leetcode.cn/problems/count-number-of-trapezoids-i/
// 难度：中等
// 题解：
// 由于本题只关注水平梯形（即梯形的点均处于 y = b 水平线上），因此需要先记录每条水平线上点的个数
// 然后遍历每条边获取任选两点的组合数，将当前边的组合数与以前边累计的组合数相乘得到水平梯形的个数
// 最后累加得到总的水平梯形个数

func countTrapezoids(points [][]int) int {
	const mod = 1_000_000_007

	// 记录水平线点的个数
	cnt := make(map[int]int)
	for _, point := range points {
		cnt[point[1]]++
	}

	ans := 0 // ans 记录水平梯形个数
	s := 0   // s   累计边数
	for _, c := range cnt {
		k := c * (c - 1) / 2 // k 表示当前行的边数 Combination(2, c)
		ans += s * k
		s += k
	}

	return ans % mod
}
