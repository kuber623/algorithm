package leetcode_342

import "math/bits"

// https://leetcode.cn/problems/power-of-four/
// 难度：简单

func isPowerOfFour(n int) bool {
	return bits.OnesCount(uint(n)) == 1 && bits.Len(uint(n))%2 == 1
}
