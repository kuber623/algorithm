package leetcode_693

// https://leetcode.cn/problems/binary-number-with-alternating-bits/
// 难度：简单

func hasAlternatingBits(n int) bool {
	u := uint(n + n>>1)
	return u&(u+1) == 0
}
