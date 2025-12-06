package leetcode_231

// https://leetcode.cn/problems/power-of-two/
// 难度：简单

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
