package leetcode_1744

// https://leetcode.cn/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/
// 难度：中等
// 题解：通过前缀和快速计算最喜欢糖果所属的颗数区间

func canEat(candiesCount []int, queries [][]int) (ans []bool) {
	n := len(candiesCount)

	// 计算前缀和
	prefsum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefsum[i+1] = prefsum[i] + candiesCount[i]
	}

	for _, query := range queries {
		ftype, fday, dcap := query[0], query[1], query[2]

		// 计算能出到的糖果数区间范围，左值为每天吃 1 颗，右值为每天吃 dcap 颗
		x1, y1 := fday+1, dcap*(fday+1)
		// 通过前缀和获取最喜欢糖果的糖果数区间
		x2, y2 := prefsum[ftype]+1, prefsum[ftype+1]
		// 求区间 [x1, y1] 与 [x2, y2] 是否存在交集
		ans = append(ans, !((y1 < x2) || (y2 < x1)))
	}

	return ans
}
