package leetcode_3226

import "math/bits"

// https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/
// 难度：简单

func minChanges(n int, k int) int {
	if n&k != k {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}
