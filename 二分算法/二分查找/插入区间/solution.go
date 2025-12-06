package leetcode_57

// https://leetcode.cn/problems/insert-interval/
// 难度：中等
// 题解：二分搜索 + 区间合并

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	if len(intervals) == 0 || newInterval[1] < intervals[0][0] {
		ans = append(ans, newInterval)
		ans = append(ans, intervals...)
		return
	}

	// 二分搜索
	// 寻找待插入区间中首个左端点大于新区间的位置
	l, r := 0, len(intervals)-1
	for l <= r {
		m := l + (r-l)>>1
		if intervals[m][0] < newInterval[0] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	intervals = append([][]int{newInterval}, intervals[l:]...)

	for i := 0; i < len(intervals); i++ {
		n := len(ans)
		if n > 0 && intervals[i][0] <= ans[n-1][1] { // 存在交集
			ans[n-1][1] = max(ans[n-1][1], intervals[i][1])
		} else { // 不存在交集
			ans = append(ans, intervals[i])
		}
	}

	return
}
