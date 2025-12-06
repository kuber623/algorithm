package leetcode_461

import "math/bits"

// https://leetcode.cn/problems/hamming-distance/
// 难度：简单

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
