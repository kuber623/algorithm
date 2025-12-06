package leetcode_3133

// https://leetcode.cn/problems/minimum-array-end/
// 难度：中等

func minEnd(n int, x int) int64 {
	mask := 0

	for k := x; k > 0; {
		b := lowbit(k)
		mask |= b
		k -= b
	}
	mask = ^mask

	for k := n - 1; k > 0; k = k >> 1 {
		b := lowbit(mask)
		mask -= b
		x += b * (k & 1)
	}

	return int64(x)
}

func lowbit(x int) int {
	return x & -x
}
